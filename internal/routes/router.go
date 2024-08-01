package routes

import (
	"gin_vue_admin_framework/internal/middleware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	//中间件
	middleware.HandleSession(Router)

	//加载路由
	LoadExampleRoutes(Router)
}
