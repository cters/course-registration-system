package initialize

import (
	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/middlewares"
	"github.com/QuanCters/backend/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.CorsMiddleware([]string{"*"}))
	r.Use(middlewares.ValidatorMiddleware())
	// r.Use(middlewares.NewRateLimiter().GlobalRateLimiter()) // 100 request / s
	// r.GET("/ping/100", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong 100",
	// 	})
	// })

	// r.Use(middlewares.NewRateLimiter().PublicAPIRateLimiter())
	// r.GET("/ping/80", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong 80",
	// 	})
	// })

	// r.Use(middlewares.NewRateLimiter().PrivateAPIRateLimiter())
	// r.GET("/ping/60", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong 60",
	// 	})
	// })

	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/checkstatus", func(c* gin.Context){c.JSON(200,gin.H{"message":"ok"})})
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}

	return r
}
