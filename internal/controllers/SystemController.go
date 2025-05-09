package controllers

import (
	"gin_vue_admin_framework/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemController struct{}

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
