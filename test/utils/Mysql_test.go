package utils

import (
	"gin_vue_admin_framework/configs"
	"gin_vue_admin_framework/utils"
	"testing"
)

func TestConnect(t *testing.T) {
	configs.InitConfig()
	utils.InitDB()
	test_db := utils.Db

	if test_db == nil {
		t.Error("db is nil")
	}
}
