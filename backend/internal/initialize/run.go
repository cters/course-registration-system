package initialize

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	var ctx = context.Background()

	Loadconfig()
	InitLogger()
	InitPostgresql()
	InitRedis(ctx)
	InitRabbitMQ(ctx)
	InitServiceInterface()
	InitConsumer()
	r := InitRouter()
	return r
}