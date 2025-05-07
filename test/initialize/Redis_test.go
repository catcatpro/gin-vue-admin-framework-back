package initialize_test

import (
	"context"
	"gin_vue_admin_framework/common"
	"gin_vue_admin_framework/initialize"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	// configs.InitConfig()
	initialize.InitSystem()
	test_rdb := common.COM_REDIS
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
