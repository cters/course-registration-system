package initialize

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)


type RedisConfig struct {
	Host string
	Port int
	Password string
	Database int
}


func InitRedis(ctx context.Context, config RedisConfig, logger *zap.Logger) (*redis.Client, error) {
	const maxRetries = 3
	var rdb *redis.Client

	for retry := 0; retry <= maxRetries; retry++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Error("Recovered from Redis panic",
					zap.Any("error", r),
					zap.Int("retry_count", retry),
					zap.Int("maxRetries", maxRetries),
					zap.String("stack", string(debug.Stack())),
				)
			}
			}()

			rdb = redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%s:%v", config.Host, config.Port),
				Password: config.Password,
				DB: config.Database,
				PoolSize: 100,
			})

			if _, err := rdb.Ping(ctx).Result(); err != nil {
				panic(fmt.Errorf("redis ping failed: %w", err))
			}
		}()
		if rdb != nil {
			if _, err := rdb.Ping(ctx).Result(); err == nil {
				logger.Info("Redis initialized successfully",)
				return rdb, nil
			}
		}

		if retry < maxRetries {
			backoff := time.Duration((retry+1)*(retry+1)) * time.Second
			logger.Warn("Retrying Redis connection...",
				zap.Int("attempt", retry+1),
				zap.Duration("backoff", backoff),
			)
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
	}

	return nil, fmt.Errorf("failed to connect to Redis after %d attempts", maxRetries)
}