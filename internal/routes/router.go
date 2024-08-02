package routes

import (
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

	}
	//middleware.HandleSession(Router)
	//store := Session.InitSession()
	//Router.Use(sessions.Sessions("systemSession", store))

	//加载路由
	{
		new(ExampleRouter).initRouter(publicRouter, privateRouter)
	}
}
