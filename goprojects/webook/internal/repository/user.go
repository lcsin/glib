package repository

import (
	"context"

	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/repository/dao"
	"github.com/lcsin/webook/internal/repository/model"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (ur *UserRepository) Create(ctx context.Context, user domain.User) error {
	return ur.dao.Insert(ctx, model.User{
		Email:    user.Email,
		Username: user.Username,
		Passwd:   user.Passwd,
		Age:      user.Age,
	})
}

func (ur *UserRepository) ModifyByID(ctx context.Context, user domain.User) error {
	return ur.dao.UpdateByID(ctx, model.User{
		ID:       user.ID,
		Username: user.Username,
		Age:      user.Age,
	})
}

func (ur *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := ur.dao.SelectByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           user.ID,
		Email:        user.Email,
		Passwd:       user.Passwd,
		Username:     user.Username,
		Age:          user.Age,
		RegisterTime: user.CreatedTime,
	}, nil
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := ur.dao.SelectByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:           user.ID,
		Email:        user.Email,
		Passwd:       user.Passwd,
		Username:     user.Username,
		Age:          user.Age,
		RegisterTime: user.CreatedTime,
	}, nil
}
