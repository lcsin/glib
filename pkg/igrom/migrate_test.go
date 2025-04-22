package igrom

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type column struct {
	ColumnName string
	ColumnType string // 跟数据库一致
	Len        string
	NotNull    bool
	Comment    string // 注释
}

func TestTable(t *testing.T) {
	dsn := "root:root@tcp(localhost:13306)/testdb1?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		t.Fatal(err)
	}

	name := "test_tbl"

	columns := []*ColumnMigrate{
		{Name: "name", Type: "varchar", Len: 255, NotNULL: true, Comment: "姓名"},
		{Name: "age", Type: "int", Comment: "年龄"},
		{Name: "score", Type: "float", Comment: "分数"},
	}

	if err = NewTableMigrate(db, name).CreateTable(columns...); err != nil {
		t.Fatal(err)
	}

	//t.Log(NewTableMigrate(db, name).HasTable())
	//t.Log(NewTableMigrate(db, name).DropColumn("score"))
	//t.Log(NewTableMigrate(db, name).AddColumn(&ColumnMigrate{
	//	Name:    "score",
	//	Type:    "float",
	//	Len:     10,
	//	Decimal: 2,
	//	Comment: "分数",
	//}))
	//t.Log(NewTableMigrate(db, name).RenameColumn("score1", "score"))
}
