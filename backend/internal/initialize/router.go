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

	r.Use(middlewares.ValidatorMiddleware())

	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/checkstatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}

	return r
}
