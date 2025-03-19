package dao

import (
	"github.com/lcsin/webook/internal/repository/model"
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{})
}
