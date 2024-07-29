package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExampleController struct{}

func (ExampleController *ExampleController) IndexAction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
