package routes

import (
	"gin_vue_admin_framework/internal/controllers"

	"github.com/gin-gonic/gin"
)

type SystemRouter struct {
}

var systemController *controllers.SystemController = new(controllers.SystemController)

func (ur *SystemRouter) InitRouter(PublicRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	// authRouter := PrivateRouter.Group("user")
	sysGroup := PublicRouter.Group("sys")
	{
		sysGroup.GET("get_captcha", systemController.GetCaptchaAction)
	}

	sysAuthGroup := PrivateRouter.Group("sys")
	{
		sysAuthGroup.POST("update_sys_settings", systemController.UpdateSysSettingsAction)
		sysAuthGroup.GET("get_sys_settings", systemController.GetSysSettingsAction)
	}

}
