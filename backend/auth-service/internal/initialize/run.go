package initialize

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() (*gin.Engine, string) {
	// 1> Read config -> environment variables
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// 2> initialize logger
	logger, err := NewLogger(LoggerConfig{
		LogLevel:    config.LogLevel,
		LogFile:     config.LogFile,
		MaxSize:     config.LogMaxSize,
		MaxBackups:  config.LogMaxBackup,
		MaxAge:      config.LogMaxAge,
		Compress:    config.LogCompress,
	})
	if err != nil {
		log.Fatalf("Could not initialize logger: %v", err)
	}

	// 3> initialize redis
	ctx, cancel :=  context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redisClient, err := InitRedis(ctx, RedisConfig{
		Host:     config.RedisHost,
		Port:     config.RedisPort,
		Password: config.RedisPassword,
		Database: config.RedisDB,
	}, logger)
	if err != nil {
		logger.Fatal("Could not initialize Redis", zap.Error(err))
	}

	// 4> Initialize database connection
	db, err := InitDB(config, logger)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// 5> Initialize router
	r := InitRouter(config, db, logger, redisClient)

	// 6> Initialize other services if needed (e.g., cache, message queue, etc.)
	return r, config.ServerPort
}
