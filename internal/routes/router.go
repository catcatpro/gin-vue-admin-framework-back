package routes

import (
	"gin_vue_admin_framework/internal/middlewares"
	AdminRoutes "gin_vue_admin_framework/internal/routes/admin"

	"github.com/gin-gonic/gin"
)

type BaseRouter interface {
	initRouter(PublicRouter *gin.Engine, PrivateRouter *gin.Engine)
}

var Router *gin.Engine

func init() {
	Router = gin.Default()
	publicRouter := Router.Group("/public")
	privateRouter := Router.Group("")
	//中间件
	{
		privateRouter.Use(middlewares.AuthRequired())
	}
	//middleware.HandleSession(Router)
	//store := Session.InitSession()
	//Router.Use(sessions.Sessions("systemSession", store))

	//加载路由
	{
		new(ExampleRouter).InitRouter(publicRouter, privateRouter)
		new(SystemRouter).InitRouter(publicRouter, privateRouter)
		new(AdminRoutes.AdminUserRouter).InitRouter(publicRouter, privateRouter)
	}
}
