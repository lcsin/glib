package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lcsin/webook/internal/domain"
	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	cmd redis.Cmdable
}

func NewUserCache(cmd redis.Cmdable) *UserCache {
	return &UserCache{
		cmd: cmd,
	}
}

func (u *UserCache) key(uid int64) string {
	return fmt.Sprintf("user:info:%d", uid)
}

func (u *UserCache) Set(ctx context.Context, user domain.User) error {
	data, _ := json.Marshal(user)
	key := u.key(user.ID)
	return u.cmd.Set(ctx, key, data, time.Hour*24).Err()
}

func (u *UserCache) Get(ctx context.Context, uid int64) (*domain.User, error) {
	key := u.key(uid)
	data, err := u.cmd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var user domain.User
	if err = json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserCache) Delete(ctx context.Context, uid int64) error {
	key := u.key(uid)
	return u.cmd.Del(ctx, key).Err()
}
