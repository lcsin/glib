package lua

import _ "embed" // embed lua script

// RedisDistributeLock Redis分布式锁脚本-获取锁
//
//go:embed redis_distribute_lock.lua
var RedisDistributeLock string

// RedisDistributeUnLock Redis分布式锁脚本-释放锁
//
//go:embed redis_distribute_unlock.lua
var RedisDistributeUnLock string

// RedisSlideWindowRateLimiter Redis基于滑动窗口算法的限流
//
//go:embed slide_window_limiter.lua
var RedisSlideWindowRateLimiter string
