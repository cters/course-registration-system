package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/shared-libs/pkg/response"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *AuthHandler) {
	// Register the routes for authentication
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/register", response.Wrap(handler.CreateAccount))
		authGroup.POST("/login", response.Wrap(handler.Login))
		authGroup.POST("/logout", response.Wrap(handler.Logout))
	}
}