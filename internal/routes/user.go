package routes

import (
	"gin_vue_admin_framework/internal/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

var userController *controllers.UserController = new(controllers.UserController)

func (ur *UserRouter) initRouter(PublicRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	// authRouter := PrivateRouter.Group("user")
	userGroup := PublicRouter.Group("user")

	{
		userGroup.POST("login", userController.LoginAction)
	}
}
