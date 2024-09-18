package controllers

import (
	"gin_vue_admin_framework/internal/models/requests"
	"gin_vue_admin_framework/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController *UserController) LoginAction(c *gin.Context) {
	loginInfo := requests.SysLoginRequest{}
	err := c.ShouldBind(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Server error occured!",
			"data":   "{}",
		})

		return
	}
	cs := services.CommmonService{}
	token, err := cs.Login(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"msg":    err.Error(),
			"data":   "{}",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "success",
		"data": gin.H{
			"token": token,
		},
	})

}
