package routes

import (
	"gin_vue_admin_framework/internal/controllers"

	"github.com/gin-gonic/gin"
)

type SystemRouter struct {
}

var systemController *controllers.SystemController = new(controllers.SystemController)

func (ur *SystemRouter) initRouter(PublicRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	// authRouter := PrivateRouter.Group("user")
	userGroup := PublicRouter.Group("sys")
	{
		userGroup.POST("login", systemController.SysLoginAction)
		userGroup.GET("get_captcha", systemController.GetCaptchaAction)
	}

}
