package initialize

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/auth/controller/http"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/middleware"
	setting "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/pkg"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/shared-libs/pkg/response"
	"go.uber.org/zap"
)

func InitRouter(
	config setting.Config,
	db *sql.DB, 
	logger *zap.Logger,
	redisClient *redis.Client,
) *gin.Engine {
	// Initialize the router
	// This function will set up the routes and middlware for the application
	// It will return a gin.Engine instance that can be used to run the server
	var r *gin.Engine
	// Set the mode based on the environment
	if config.LogLevel == "debug" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	r.Use(middleware.CORS) // cross
	r.Use(middleware.ValidatorMiddleware())
	r.Use(middleware.LoggerMiddleware(logger)) // logging

	// r.Use() // limiter global
	// r.Use(middlewares.Validator())      // middleware

	// r.Use(middlewares.NewRateLimiter().GlobalRateLimiter()) // 100 req/s
	r.GET("/ping/100", func(ctx *gin.Context) {

		response.SuccessResponse(ctx, 200, "pong", nil)
	})

	r.GET("/ping/200", response.Wrap(func(ctx *gin.Context) (res interface{}, err error) {
		return "pong", nil
	}))

	// === Đăng ký routes theo module
	v1 := r.Group("/v1/api")
	// Register the auth routes
	// === DI các handler
	authHandler := InitAuth(
		config,
		db, 
		logger,
		redisClient,	
	)
	http.RegisterAuthRoutes(v1, authHandler)

	return r

}