package user

import (
	"github.com/QuanCters/backend/internal/controller/user"
	"github.com/QuanCters/backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}
func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/login", user.UserLoginController.Login)
	}
	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthenMiddleware())
	{
		userRouterPublic.POST("/register", user.UserAdminController.Register)
	}
}