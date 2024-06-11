package mysql

import (
	"gin_vue_admin_framework/pkg/Db/mysql"
	"testing"
)

func TestConnect(t *testing.T) {

	test_db := mysql.Db

	if test_db == nil {
		t.Error("db is nil")
	}
}
