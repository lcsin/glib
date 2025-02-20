package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// SimpleTask 简单的定时任务，不建议在生产环境中使用
func SimpleTask(duration time.Duration, callback func()) {
	go func() {
		t := time.NewTicker(duration)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				callback()
			}
		}
	}()
}

// CronTask Cron定时任务
func CronTask(spec string, callback func()) {
	c := cron.New(cron.WithSeconds())

	if _, err := c.AddFunc(spec, callback); err != nil {
		log.Printf("add corn task error: %v", err)
		return
	}

	c.Start()
}

// TimeoutTask 带过期时间的任务
func TimeoutTask(ctx context.Context, timeout time.Duration, fn func(timeoutCtx context.Context) error) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- fn(timeoutCtx)
	}()

	select {
	case <-timeoutCtx.Done():
		return fmt.Errorf("任务超时")
	case err := <-done:
		return err
	}
}
