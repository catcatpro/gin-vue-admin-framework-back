package services

import (
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

func (ss SysService) SystemSettingsSave(req *[]requests.SystemSettingsRequest) error {
	var list []system.SystemSettingsModel
	for _, item := range *req {
		list = append(list, system.SystemSettingsModel{
			SetKey:   item.SetKey,
			SetValue: item.SetValue,
		})
	}

	var systemSettings system.SystemSettingsModel
	return systemSettings.SaveSysSettings(list)
}
