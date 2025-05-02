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
	InitServiceInterface()
	InitRedis(ctx)
	InitRabbitMQ(ctx)
	r := InitRouter()
	return r
}