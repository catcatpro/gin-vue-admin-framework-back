package utils

import (
	"context"
	"gin_vue_admin_framework/configs"
	"gin_vue_admin_framework/utils"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	configs.InitConfig()
	j := utils.NewJWT()
	claims := j.CreateClaims(1, "user1")
	token, err := j.CreateToken(claims)
	if err != nil {
		t.Error(err)
	}
	claims2, err := j.ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	if claims2.Username != "user1" {
		t.Error("should be user1 for parse token")
	}
	//t.Error("test")
}

/**
* 单元测试
* 将token存储在redis
 */
func TestJwtStoreForRedis(t *testing.T) {
	configs.InitConfig()
	j := utils.NewJWT()

	claims := j.CreateClaims(1, "user1")
	token, err := j.CreateToken(claims)
	if err != nil {
		t.Error(err)
	}

	utils.InitRedis()
	test_rdb := utils.Rdb
	if test_rdb == nil {
		t.Error("rdb is nil")
	}

	ctx := context.Background()
	_, err = test_rdb.Set(ctx, "1", token, time.Duration(30)*time.Second).Result()
	if err != nil {
		t.Error(err)
	}
	val, err := test_rdb.Get(ctx, "1").Result()
	if err != nil {
		t.Error(err)
	}
	if val != token {
		t.Error("token is wrong")
	}
}
