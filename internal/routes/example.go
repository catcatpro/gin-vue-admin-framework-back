package routes

import (
	"gin_vue_admin_framework/internal/controllers"
	"github.com/gin-gonic/gin"
)

var exampleController *controllers.ExampleController = new(controllers.ExampleController)

func LoadExampleRoutes(r *gin.Engine) *gin.RouterGroup {
	exampleGroup := r.Group("/example")
	{
		exampleGroup.GET("/index", exampleController.IndexAction)
	}
	return exampleGroup
}
