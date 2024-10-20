package utils

import (
	"errors"
	"image/color"
	"time"

	"github.com/mojocn/base64Captcha"
	"golang.org/x/exp/rand"
)

func init() {
	//init rand seed
	rand.Seed(uint64(time.Now().UnixNano()))
}

var store = base64Captcha.DefaultMemStore

/**
* 接口
* 系统生成和验证验证码方法
* 使用Redis作为缓存：id:生成唯一id，value为值，有效期5分钟
 */
type CaptchaInterfaceV2 interface {
	Generate() (id, base64str string, err error)
	Verify(id string, awaer string) (res bool, err error)
}

type CaptchaV2 struct {
}

var stringCaptchaDriver base64Captcha.Driver = &base64Captcha.DriverString{
	Source:          "1234567890",
	Width:           100,
	Height:          50,
	NoiseCount:      0,
	ShowLineOptions: base64Captcha.OptionShowHollowLine,
	Length:          5,
	BgColor: &color.RGBA{
		R: 40,
		G: 30,
		B: 29,
		A: 80,
	},
	Fonts: nil,
}

// 缓存验证码
func (c2 CaptchaV2) Generate() (id, base64str string, err error) {
	var driver base64Captcha.Driver = stringCaptchaDriver
	c := base64Captcha.NewCaptcha(driver, store)
	id, base64str, _, err = c.Generate()
	if err != nil {
		return "", "", err
	}
	return
}

// 验证验证码
func (c CaptchaV2) Verify(id string, awaer string) (res bool, err error) {
	res = store.Verify(id, awaer, false)
	if !res {
		err = errors.New("Incorrect Captcha")
	}
	return
}
