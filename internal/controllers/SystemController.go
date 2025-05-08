package controllers

import (
	"fmt"
	"gin_vue_admin_framework/internal/models/requests"

	"gin_vue_admin_framework/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemController struct{}

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

func (sc *SystemController) GetCaptchaAction(c *gin.Context) {
	ss := services.SysService{}
	s_captcha, err := ss.GenerateCaptcha()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
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
			"id":      s_captcha.Id,
			"captcha": s_captcha.Captcha,
		},
	})

}
