package repository

import (
	"context"

	"ginlayout/internal/domain"
	"ginlayout/internal/repository/cache"
	"ginlayout/internal/repository/dao"
)

type IHelloRepository interface {
	Get(ctx context.Context, id int64) (*domain.Hello, error)
}

type HelloRepository struct {
	dao   dao.IHelloDAO
	cache cache.IHelloCache
}

func NewHelloRepository(dao dao.IHelloDAO, cache cache.IHelloCache) IHelloRepository {
	return &HelloRepository{dao: dao, cache: cache}
}

func (h HelloRepository) Get(ctx context.Context, id int64) (*domain.Hello, error) {
	// 先从缓存中查询
	hello, err := h.cache.Get(ctx, id)
	if err == nil {
		return hello, nil
	}

	// 从数据库中查询
	_, err = h.dao.SelectByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Hello{
		// 填充 model.Hello
	}, nil
}
