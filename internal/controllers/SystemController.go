package controllers

import (
	"gin_vue_admin_framework/internal/services"
	"gin_vue_admin_framework/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemController struct{}

func (sc *SystemController) GetCaptchaAction(c *gin.Context) {
	ss := services.SysService{}
	s_captcha, err := ss.GenerateCaptcha()
	if err != nil {
		//c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		//	"status": "error",
		//	"msg":    err.Error(),
		//	"data":   gin.H{},
		//})
		utils.FailResponse(c, http.StatusInternalServerError, err.Error(), gin.H{})
		return
	}

	utils.SuccessResponse(c, gin.H{
		"id":      s_captcha.Id,
		"captcha": s_captcha.Captcha,
	})

}
