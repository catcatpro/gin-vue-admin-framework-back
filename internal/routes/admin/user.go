package AdminRoutes

import (
	adminControllers "gin_vue_admin_framework/internal/controllers/admin"
	"github.com/gin-gonic/gin"
)

type AdminUserRouter struct{}

var adminUserController = new(adminControllers.UserController)

func (ur *AdminUserRouter) InitRouter(PublicRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	adminUserGroup := PublicRouter.Group("admin_user")
	{
		adminUserGroup.POST("login", adminUserController.SysLoginAction)
	}
}
