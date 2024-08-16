package mysql

import (
	"gin_vue_admin_framework/cmd"
	"gin_vue_admin_framework/utils"
	"testing"
)

func TestConnect(t *testing.T) {
	cmd.InitSystem()
	test_db := utils.Db

	if test_db == nil {
		t.Error("db is nil")
	}
}
