package utils

import (
	"context"
	"errors"
	"time"
)

/**
* 接口
* 系统生成和验证验证码方法
* 使用Redis作为缓存：id:生成唯一id，value为值，有效期5分钟
 */
type CaptchaInterface interface {
	Set(id string, value string) error
	Verify(id string, awaer string) (bool, error)
}

type Captcha struct{}

// 缓存验证码
func (c Captcha) Set(id string, value string) error {
	ctx := context.Background()
	_, err := Rdb.Set(ctx, id, value, time.Duration(300)*time.Second).Result()
	if err != nil {
		return err
	}

	return nil
}

// 验证验证码
func (c Captcha) Verify(id string, awaer string) (bool, error) {
	ctx := context.Background()
	res, err := Rdb.Get(ctx, id).Result()

	if err != nil || res != awaer {
		return false, errors.New("Incorrect Captcha")
	}
	return true, nil
}
