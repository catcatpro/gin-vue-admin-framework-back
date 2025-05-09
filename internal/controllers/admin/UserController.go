package controllers

import (
	"fmt"
	"gin_vue_admin_framework/internal/models/requests"
	"gin_vue_admin_framework/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (sc *SystemController) SysLoginAction(c *gin.Context) {
	loginInfo := requests.SysLoginRequest{}
	err := c.ShouldBind(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Parameter error",
			"data":   "{}",
		})

		return
	}
	fmt.Println(loginInfo)
	cs := services.SysUserService{}
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
