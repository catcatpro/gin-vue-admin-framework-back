package services

import (
	"errors"
	"gin_vue_admin_framework/internal/models"
	"gin_vue_admin_framework/internal/models/requests"
)

type CommmonService struct{}

func (cs *CommmonService) Login(UserReq *requests.LoginRequest) (bool, error) {

	//TODO 临时验证码验证
	if UserReq.Captcha != "2333" {
		return false, errors.New("Incorrect Captcha")
	}

	var userInfo models.User
	userInfo.Username = UserReq.Username
	userInfo.Password = UserReq.Password

	//执行登录验证
	if !userInfo.UserLogin() {
		return false, errors.New("Username or Password incorrect")
	}
	return true, nil
}
