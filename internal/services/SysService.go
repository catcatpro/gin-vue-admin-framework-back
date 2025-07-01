package services

import (
	"gin_vue_admin_framework/internal/models"
	"gin_vue_admin_framework/internal/models/requests"
	"gin_vue_admin_framework/internal/models/system"
	"gin_vue_admin_framework/utils"
)

type SysService struct {
}

func (ss SysService) GenerateCaptcha() (sc system.SysCaptchaModel, err error) {
	var s_captcha system.SysCaptchaModel

	var cap utils.CaptchaInterfaceV2
	cap = new(utils.CaptchaV2)
	id, captch, err := cap.Generate()
	if err != nil {
		return s_captcha, err
	}

	s_captcha.Id = id
	s_captcha.Captcha = captch
	sc = s_captcha
	return

}

func (ss SysService) SystemSettingsUpdate(req *[]requests.SystemSettingsRequest) error {
	var list []models.SystemSettings
	for _, item := range *req {
		list = append(list, models.SystemSettings{
			SetKey:   item.SetKey,
			SetValue: item.SetValue,
		})
	}

	var systemSettings models.SystemSettings
	return systemSettings.UpdateSysSettings(list)
}
