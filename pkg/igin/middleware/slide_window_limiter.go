package middleware

import (
	_ "embed"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

//go:embed slide_window_limiter.lua
var luaScript string

type SlideWindowLimiterBuilder struct {
	cmd redis.Cmdable

	key       string        // 限流对象
	interval  time.Duration // 限流间隔
	threshold int           // 限流阈值
}

func NewSlideWindowLimiterBuilder(cmd redis.Cmdable, key string, interval time.Duration, threshold int) *SlideWindowLimiterBuilder {
	return &SlideWindowLimiterBuilder{
		cmd:       cmd,
		key:       key,
		interval:  interval,
		threshold: threshold,
	}
}

func (b *SlideWindowLimiterBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		limited, err := b.limit(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "internal server error",
			})
			return
		}
		if limited {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "too many requests",
			})
			return
		}
		c.Next()
	}
}

func (b *SlideWindowLimiterBuilder) limit(ctx *gin.Context) (bool, error) {
	return b.cmd.Eval(ctx, luaScript, []string{b.key},
		b.interval.Milliseconds(), b.threshold, time.Now().UnixMilli()).Bool()
}
