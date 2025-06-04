package utils

import (
	"errors"
	"fmt"
	"gin_vue_admin_framework/configs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret []byte
}

type CustomClaims struct {
	Username string `json:"username"`
	Id       uint   `json:"id"`
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		secret: []byte(configs.SystemConfigs.Jwt.Secret),
	}
}

func (j *JWT) CreateClaims(id uint, username string) *CustomClaims {
	day, _ := strconv.ParseInt(configs.SystemConfigs.Jwt.ExpiresAt, 10, 64)
	return &CustomClaims{
		username,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(day) * time.Hour)), //day小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    configs.SystemConfigs.Jwt.Issuer,
		},
	}
}

// 生成jwt.token
func (j *JWT) CreateToken(claims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

// 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil && !token.Valid {
		err = errors.New("invalid token")
	}
	exp_time := claims.ExpiresAt
	fmt.Println("exp_time", exp_time)
	return claims, err
}

// 验证过期时间
func (j *JWT) VerifyTokenExpiresAt(tokenString string) (*CustomClaims, error) {
	claims, err := j.ParseToken(tokenString)
	if err != nil {
		err = errors.New("invalid token")
	}
	if claims.ExpiresAt.Unix() <= time.Now().Unix() {
		err = errors.New("token is expired")
	}
	return claims, err
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	_, err := j.ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	claims := &CustomClaims{}
	claims, err = j.ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	nowClaims := j.CreateClaims(claims.Id, claims.Username)
	return j.CreateToken(nowClaims)

}

func (j *JWT) removeToken(username string) (string, error) {}
