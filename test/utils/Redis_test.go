package utils

import (
	"context"
	"gin_vue_admin_framework/configs"
	"gin_vue_admin_framework/utils"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	configs.InitConfig()
	utils.InitRedis()
	test_rdb := utils.Rdb
	if test_rdb == nil {
		t.Error("rdb is nil")
	}

	ctx := context.Background()
	_, err := test_rdb.Set(ctx, "key", "val", time.Duration(30)*time.Second).Result()
	if err != nil {
		t.Error(err)
	}

	val, err := test_rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
	}
	if val != "val" {
		t.Error("val != val")
	}
}
