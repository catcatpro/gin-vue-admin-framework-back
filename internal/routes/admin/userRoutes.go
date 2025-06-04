package AdminRoutes

import (
	adminControllers "gin_vue_admin_framework/internal/controllers/admin"

	"github.com/gin-gonic/gin"
)

type AdminUserRoutes struct{}

var adminUserController = new(adminControllers.UserController)

func (ur *AdminUserRoutes) InitRouter(PublicRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	adminUserGroup := PublicRouter.Group("admin/user")
	{
		adminUserGroup.POST("login", adminUserController.LoginAction)
		adminUserGroup.POST("test_create", adminUserController.CreateUserAction)
	}

	adminUserPrivateGroup := PrivateRouter.Group("admin/user")
	{
		adminUserPrivateGroup.POST("create", adminUserController.CreateUserAction)
		adminUserPrivateGroup.POST("get_user_info", adminUserController.GetUserInfoAction)
	}
}
