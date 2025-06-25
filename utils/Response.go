package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "success",
		"data":   data,
	})
}

func FailResponse(c *gin.Context, status_code int, msg string, data interface{}) {
	c.AbortWithStatusJSON(status_code, gin.H{
		"status": "error",
		"msg":    msg,
		"data":   data,
	})
}
