package mysql

import (
	"gin_vue_admin_framework/utils/gorm"
	"testing"
)

func TestConnect(t *testing.T) {

	test_db := gorm.Db

	if test_db == nil {
		t.Error("db is nil")
	}
}
