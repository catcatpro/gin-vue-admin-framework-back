package controllers

import (
	"fmt"
	"gin_vue_admin_framework/internal/models/requests"
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

// 获取系统设置
func (sc *SystemController) GetSysSettingsAction(c *gin.Context) {
	sysSettings, err := services.SysService{}.GetSystemSettings()
	if err != nil {
		fmt.Println("err2", err)
		utils.FailResponse(c, http.StatusInternalServerError, "Server Error: ", "")
		return
	}

	utils.SuccessResponse(c, sysSettings)
}

// 更新系统设置
func (controller *SystemController) UpdateSysSettingsAction(c *gin.Context) {
	var err error
	list := []requests.SystemSettingsRequest{}
	err = c.ShouldBindBodyWithJSON(&list)
	if err != nil {
		utils.FailResponse(c, http.StatusBadRequest, "Parameter error", "")
		return
	}
	fmt.Println("list", list)
	err = services.SysService{}.SystemSettingsUpdate(&list)
	fmt.Println("err", err)
	if err != nil {
		utils.FailResponse(c, http.StatusInternalServerError, "Server Error", "")
		return
	}
	utils.SuccessResponse(c, gin.H{})
}
