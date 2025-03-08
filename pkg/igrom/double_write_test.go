package igrom

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserTbl struct {
	ID   int64 `gorm:"primaryKey;autoIncrement"`
	Name string
}

func (u *UserTbl) TableName() string {
	return "user_tbl"
}

func TestDoubleWrite(t *testing.T) {
	dnsSrc := "root:root@tcp(localhost:13306)/testdb1?charset=utf8mb4&parseTime=True"
	dbSrc, err := gorm.Open(mysql.Open(dnsSrc))
	require.NoError(t, err)
	_ = dbSrc.AutoMigrate(&UserTbl{})

	dnsDst := "root:root@tcp(localhost:13306)/testdb2?charset=utf8mb4&parseTime=True"
	dbDst, err := gorm.Open(mysql.Open(dnsDst))
	require.NoError(t, err)
	_ = dbDst.AutoMigrate(&UserTbl{})

	var pattern atomic.Value
	pattern.Store(PatternSrcFirst)
	dbDw, err := gorm.Open(mysql.New(mysql.Config{
		Conn: &DoubleWritePool{
			src:     dbSrc.ConnPool,
			dst:     dbDst.ConnPool,
			pattern: pattern,
		},
	}))
	require.NoError(t, err)
	dbDw.Create(&UserTbl{
		Name: "zhangsan",
	})

	err = dbDw.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&UserTbl{
			Name: "zhangsan_tx",
		}).Error; err != nil {
			return err
		}

		// 测试回滚
		//if err = tx.Model(&UserTbl{}).Where("id = ?", 2).UpdateColumn("name", "zhangsan_tx_update").Error; err != nil {
		//	return err
		//}
		//return errors.New("test tx error")
		return nil
	})
	require.NoError(t, err)
}
