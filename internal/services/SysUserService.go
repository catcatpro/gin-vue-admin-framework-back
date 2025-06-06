package services

import (
	"errors"
	"gin_vue_admin_framework/internal/models"
	"gin_vue_admin_framework/internal/models/requests"
	"gin_vue_admin_framework/utils"
	"github.com/gin-gonic/gin"
)

type SysUserService struct{}

func (cs *SysUserService) Login(UserReq *requests.LoginRequest) (string, string, error) {
	var cap utils.CaptchaInterfaceV2
	cap = new(utils.CaptchaV2)
	//验证码验证
	if res, err := cap.Verify(UserReq.CaptchaId, UserReq.Captcha); (err != nil || !res) && gin.Mode() == gin.ReleaseMode {

		return "", "", err
	}

	var userInfo models.User
	userInfo.Username = UserReq.Username
	userInfo.Password = UserReq.Password

	//执行登录验证
	if err := userInfo.UserLogin(); err != nil {
		return "", "", errors.New("Username or Password incorrect")
	}

	j := utils.NewJWT()

	claims := j.CreateClaims(userInfo.ID, userInfo.Username)
	token, err := j.CreateToken(claims, false)

	if err != nil {
		return "", "", err
	}

	refresh_claims := j.CreateRefreshClaims(userInfo.ID)
	refresh_token, err := j.CreateToken(refresh_claims, true)
	if err != nil {
		return "", "", err
	}
	return token, refresh_token, nil
}

func (cs *SysUserService) RefreshToken(RefreshReq *requests.RefreshTokenRequest) (string, error) {
	j := utils.NewJWT()
	claims, err := j.ParseToken(RefreshReq.RefreshToken, true)
	if err != nil {
		return "", err
	}
	//验证token是否过期？
	_, err = j.VerifyTokenExpiresAt(RefreshReq.RefreshToken, true)
	if err != nil {
		return "", err
	}

	nowClaims := j.CreateRefreshClaims(claims.Id)
	now_token, err := j.CreateToken(nowClaims, false)
	if err != nil {
		return "", err
	}
	return now_token, nil
}

func (cs *SysUserService) CreateUser(req *requests.CreateUserRequest) error {

	//确认密码是否正确
	if req.ConfirmPassword != req.Password {
		return errors.New("confirm password not match")
	}

	var userInfo models.User
	userInfo.Username = req.Username
	userInfo.Password = req.Password
	if _, err := userInfo.CreateUser(); err != nil {
		return err
	}

	return nil
}

func (cs *SysUserService) GetUserInfo(req *requests.GetUserInfoRequest) (models.User, error) {
	var err error = nil
	var userInfo models.User

	switch req.Type {
	case "token":
		err = userInfo.GetUserInfoByToken(req.Data)
		break
	}
	return userInfo, err
}
