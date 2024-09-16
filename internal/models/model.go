package models

import (
	"gin_vue_admin_framework/utils"
	"time"
)

type Basic struct {
	ID          uint      `json:"id"`
	CreatedTime time.Time `json:"created_at"`
	UpdatedTime time.Time `json:"updated_at"`
}

var db = utils.Db
