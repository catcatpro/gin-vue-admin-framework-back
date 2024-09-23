package services

import (
	"gin_vue_admin_framework/internal/models/system"
	"gin_vue_admin_framework/utils"
)

type SysService struct {
}

func (ss SysService) GenerateCaptcha() (sc system.SysCaptcha, err error) {
	var s_captcha system.SysCaptcha

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
