package lua

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

func InitRedis() redis.Cmdable {
	addr := "192.168.1.103:6379"
	passwd := ""
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
	})
}

func TestLuaScriptLock(t *testing.T) {
	rdb := InitRedis()
	res, err := rdb.Eval(context.Background(), RedisDistributeLock, []string{"keys1"}, "uniqueId", 5*60).Int()
	if err != nil {
		panic(err)
	}
	t.Log("res => ", res)
}

func TestLuaScriptUnlock(t *testing.T) {
	rdb := InitRedis()
	res, err := rdb.Eval(context.Background(), RedisDistributeUnLock, []string{"keys1"}, "uniqueId").Int()
	if err != nil {
		panic(err)
	}
	t.Log("res => ", res)
}
