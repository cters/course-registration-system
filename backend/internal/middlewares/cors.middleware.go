package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(allowedOrigins []string) gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{"GET","POST","PUT","DELETE"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge: 12*time.Hour,
	}

	if len(allowedOrigins) == 1 && allowedOrigins[0] == "*" {
		config.AllowAllOrigins = true
		config.AllowOrigins = nil
	}

	return cors.New(config)
}