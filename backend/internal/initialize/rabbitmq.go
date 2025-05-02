package initialize

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/QuanCters/backend/global"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

var (
	rabbitRetryCount = 0
	rabbitMaxRetries = 3
	rabbitMutex      sync.Mutex
)

func InitRabbitMQ(ctx context.Context) {
	cfg := global.Config.RabbitMQ

	for rabbitRetryCount = 0; rabbitRetryCount <= rabbitMaxRetries; rabbitRetryCount++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					global.Logger.Error("Recovered from RabbitMQ panic",
						zap.Any("error", r),
						zap.Int("retry_count", rabbitRetryCount),
						zap.Int("maxRetries", rabbitMaxRetries),
						zap.String("stack", string(debug.Stack())),
					)
				}
			}()

			rabbitMutex.Lock()
			defer rabbitMutex.Unlock()

			connURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", 
				cfg.User, 
				cfg.Password, 
				cfg.Host, 
				cfg.Port,
			)

			conn, err := amqp091.Dial(connURL)
			if err != nil {
				panic(fmt.Sprintf("Connection failed: %v", err))
			}

			// Verify connection
			ch, err := conn.Channel()
			if err != nil {
				panic(fmt.Sprintf("Channel creation failed: %v", err))
			}
			defer ch.Close()

			global.Logger.Info("Initializing RabbitMQ successfully")
			global.Rmq = conn
			rabbitRetryCount = 0 // Reset counter on success
		}()

		if global.Rmq != nil {
			break
		}

		if rabbitRetryCount < rabbitMaxRetries {
			backoff := time.Duration((rabbitRetryCount+1)*(rabbitRetryCount+1)) * time.Second
			global.Logger.Warn("Retrying RabbitMQ connection...",
				zap.Int("attempt", rabbitRetryCount+1),
				zap.Duration("backoff", backoff),
			)
			time.Sleep(backoff)
		} else {
			global.Logger.Fatal("RabbitMQ connection failed after max retries")
		}
	}
}

