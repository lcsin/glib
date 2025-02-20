package pkg

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeoutTask(t *testing.T) {
	job := func(timeoutCtx context.Context) error {
		var counter int
		for {
			select {
			case <-timeoutCtx.Done(): // 任务超时
				return fmt.Errorf("任务超时")
			default: // 轮询执行任务
				time.Sleep(time.Second)
				counter++
				t.Logf("do job counter++: %v", counter)
				// 任务执行完成
				if counter > 20 {
					t.Log("job done!!!")
					return nil
				}
			}
		}
	}

	now := time.Now()
	if err := TimeoutTask(context.Background(), time.Second*10, job); err != nil {
		t.Log(err)
	}
	t.Log(time.Now().Sub(now))
}
