package services

import (
	"errors"
	"gin_vue_admin_framework/internal/models"
	"gin_vue_admin_framework/internal/models/requests"
	"gin_vue_admin_framework/utils"
)

type CommmonService struct{}

func (cs *CommmonService) Login(UserReq *requests.SysLoginRequest) (string, error) {
	var cap utils.CaptchaInterface
	cap = new(utils.Captcha)
	//TODO 验证码验证
	if res, err := cap.Verify(UserReq.CaptchaId, UserReq.Captcha); err != nil || !res {
		return "", err
	}

	var userInfo models.User
	userInfo.Username = UserReq.Username
	userInfo.Password = UserReq.Password

	//执行登录验证
	if !userInfo.UserLogin() {
		return "", errors.New("Username or Password incorrect")
	}

	j := utils.NewJWT()

	//TODO userInfo.ID
	claims := j.CreateClaims(1, userInfo.Username)
	token, err := j.CreateToken(claims)

	if err != nil {
		return "", err
	}
	return token, nil
}
