package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExampleController struct{}

func (ExampleController *ExampleController) IndexAction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func (ExampleController *ExampleController) TestSessionAction(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Set("test", "test_value")
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"test": session.Get("test"),
	})
}
