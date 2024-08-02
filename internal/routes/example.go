package routes

import (
	"gin_vue_admin_framework/internal/controllers"
	"github.com/gin-gonic/gin"
)

type ExampleRouter struct {
}

var exampleController *controllers.ExampleController = new(controllers.ExampleController)

func (e *ExampleRouter) initRouter(PublicRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	authRouter := PrivateRouter.Group("example")
	exampleGroup := PublicRouter.Group("example")
	{
		exampleGroup.GET("/index", exampleController.IndexAction)
		exampleGroup.GET("/test_session", exampleController.TestSessionAction)
	}
	{
		authRouter.GET("/index", exampleController.IndexAction)
	}

}
