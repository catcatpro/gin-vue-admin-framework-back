package routes

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	//加载路由
	LoadExampleRoutes(Router)
}
