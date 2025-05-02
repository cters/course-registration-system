package initialize

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)


var (
	redisRetryCount = 0
	maxRetries= 3
	redisMutex sync.Mutex // avoid race condition
)

func InitRedis(ctx context.Context) {
	r := global.Config.Redis

	for redisRetryCount = 0; redisRetryCount <= maxRetries; redisRetryCount++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					global.Logger.Error("Recovered from Redis panic",
					zap.Any("error", r),
					zap.Int("retry_count", redisRetryCount),
					zap.Int("maxRetries", maxRetries),
					zap.String("stack", string(debug.Stack())),
				)
			}
			}()

			rdb := redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%s:%v", r.Host, r.Port),
				Password: r.Password,
				DB: r.Database,
				PoolSize: 10,
			})

			_, err := rdb.Ping(ctx).Result()
			if err != nil {
				panic(err)
			}

			global.Logger.Info("Initializing Redis Successfully")
			global.Rdb = rdb
			redisRetryCount = 0
		}()

		if global.Rdb != nil {
			break
		}

		if redisRetryCount < maxRetries {
			backoff := time.Duration((redisRetryCount+1)*(redisRetryCount+1)) * time.Second
			fmt.Println(">>>>>>backoff: ", backoff)
			global.Logger.Warn("Retrying Redis connection...",
				zap.Int("attempt", redisRetryCount+1),
				zap.Duration("backoff", backoff),
			)
			time.Sleep(backoff)
		} else {
			global.Logger.Fatal("Redis connection failed after max retries")
		}
	}
}