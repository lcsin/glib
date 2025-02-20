package igrom

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Upset 插入或更新
func Upset(db *gorm.DB, updates map[string]any, create any) error {
	return db.Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(updates),
	}).Create(create).Error
}

// Paginate 分页
func Paginate(db *gorm.DB, pageNo, pageSize int64) *gorm.DB {
	if pageNo > 0 && pageSize > 0 {
		return db.Limit(int(pageSize)).Offset(int((pageNo - 1) * pageSize))
	}
	return db
}
