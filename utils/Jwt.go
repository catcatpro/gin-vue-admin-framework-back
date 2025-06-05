package utils

import (
	"errors"
	"fmt"
	"gin_vue_admin_framework/configs"
	"github.com/google/uuid"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret        []byte
	refreshSecret []byte
}

type CustomClaims struct {
	Username string `json:"username"`
	Id       uint   `json:"id"`
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		secret:        []byte(configs.SystemConfigs.Jwt.Secret),
		refreshSecret: []byte(configs.SystemConfigs.Jwt.RefreshSecret),
	}
}

func (j *JWT) CreateClaims(id uint, username string) *CustomClaims {
	expiresAt, _ := strconv.ParseInt(configs.SystemConfigs.Jwt.ExpiresAt, 10, 64)
	return &CustomClaims{
		username,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresAt) * time.Minute)), //day小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    configs.SystemConfigs.Jwt.Issuer,
		},
	}
}

func (j *JWT) CreateRefreshClaims(id uint) *CustomClaims {
	expiresAt, _ := strconv.ParseInt(configs.SystemConfigs.Jwt.RefreshExpiresAt, 10, 64)
	return &CustomClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresAt) * 60 * time.Hour)), //day小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    configs.SystemConfigs.Jwt.Issuer,
			ID:        uuid.NewString(),
		},
	}
}

// 生成jwt.token
func (j *JWT) CreateToken(claims *CustomClaims, isRefresh bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var secret []byte = j.secret
	if isRefresh {
		secret = j.refreshSecret
	}
	return token.SignedString(secret)
}

// 解析token
func (j *JWT) ParseToken(tokenString string, isRefresh bool) (*CustomClaims, error) {
	claims := &CustomClaims{}

	var secret []byte = j.secret
	if isRefresh {
		secret = j.refreshSecret
	}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil && !token.Valid {
		err = errors.New("invalid token")
	}
	exp_time := claims.ExpiresAt
	fmt.Println("exp_time", exp_time)
	return claims, err
}

// 验证过期时间
func (j *JWT) VerifyTokenExpiresAt(tokenString string, isRefresh bool) (*CustomClaims, error) {
	claims, err := j.ParseToken(tokenString, isRefresh)
	if err != nil {
		err = errors.New("invalid token")
	}
	if claims.ExpiresAt.Unix() <= time.Now().Unix() {
		err = errors.New("token is expired")
	}
	return claims, err
}

//func (j *JWT) RefreshToken(tokenString string) (string, error) {
//
//}

//func (j *JWT) removeToken(username string) (string, error) {}
