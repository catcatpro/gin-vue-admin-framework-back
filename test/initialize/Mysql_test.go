package initialize_test

import (
	"gin_vue_admin_framework/common"
	"gin_vue_admin_framework/initialize"
	"testing"
)

func TestConnect(t *testing.T) {
	// configs.InitConfig()
	initialize.InitSystem()
	test_db := common.COM_DB

	if test_db == nil {
		t.Error("db is nil")
	}
}
