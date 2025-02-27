package cache

import (
	"context"

	"ginlayout/internal/domain"

	"github.com/redis/go-redis/v9"
)

type IHelloCache interface {
	Get(ctx context.Context, id int64) (*domain.Hello, error)
}

type HelloCache struct {
	rdb redis.Cmdable
}

func NewHelloCache(rdb redis.Cmdable) IHelloCache {
	return &HelloCache{rdb: rdb}
}

func (h HelloCache) Get(ctx context.Context, id int64) (*domain.Hello, error) {
	return &domain.Hello{}, nil
}
