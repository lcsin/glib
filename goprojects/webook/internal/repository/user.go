package repository

import (
	"context"

	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/repository/cache"
	"github.com/lcsin/webook/internal/repository/dao"
	"github.com/lcsin/webook/internal/repository/model"
)

type IUserRepository interface {
	Create(ctx context.Context, user domain.User) error
	ModifyByID(ctx context.Context, user domain.User) error
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

type UserRepository struct {
	dao   dao.IUserDAO
	cache cache.IUserCache
}

func NewUserRepository(dao dao.IUserDAO, cache cache.IUserCache) IUserRepository {
	return &UserRepository{dao: dao, cache: cache}
}

func (ur *UserRepository) Create(ctx context.Context, user domain.User) error {
	defer ur.cache.Set(ctx, user)
	return ur.dao.Insert(ctx, model.User{
		Email:    user.Email,
		Username: user.Username,
		Passwd:   user.Passwd,
		Age:      user.Age,
	})
}

func (ur *UserRepository) ModifyByID(ctx context.Context, user domain.User) error {
	defer ur.cache.Delete(ctx, user.ID)
	return ur.dao.UpdateByID(ctx, model.User{
		ID:       user.ID,
		Username: user.Username,
		Age:      user.Age,
	})
}

func (ur *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	// 先从缓存中取
	cacheUser, err := ur.cache.Get(ctx, id)
	if err == nil {
		return cacheUser, nil
	}

	// 缓存中没取到，从数据库中取
	// 这里需要注意，如果缓存崩了，可能到导致大量的请求把数据库打崩
	dbUser, err := ur.dao.SelectByID(ctx, id)
	if err != nil {
		return nil, err
	}
	domainUser := domain.User{
		ID:           dbUser.ID,
		Email:        dbUser.Email,
		Passwd:       dbUser.Passwd,
		Username:     dbUser.Username,
		Age:          dbUser.Age,
		RegisterTime: dbUser.CreatedTime,
	}

	// 数据库中取到后回写缓存
	_ = ur.cache.Set(ctx, domainUser)
	return &domainUser, nil
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
