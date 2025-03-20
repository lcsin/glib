package dao

import (
	"context"
	"time"

	"github.com/lcsin/webook/internal/repository/model"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (ud *UserDAO) Insert(ctx context.Context, user model.User) error {
	user.CreatedTime = time.Now().UnixMilli()
	return ud.db.WithContext(ctx).Create(&user).Error
}

func (ud *UserDAO) UpdateByID(ctx context.Context, user model.User) error {
	return ud.db.Model(&model.User{}).WithContext(ctx).Where("id = ?", user.ID).UpdateColumns(map[string]interface{}{
		"username":     user.Username,
		"age":          user.Age,
		"updated_time": time.Now().UnixMilli(),
	}).Error
}

func (ud *UserDAO) SelectByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	if err := ud.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ud *UserDAO) SelectByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := ud.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
