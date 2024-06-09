package models

import "time"

type Basic struct {
	ID          uint      `json:"id"`
	CreatedTime time.Time `json:"created_at"`
	UpdatedTime time.Time `json:"updated_at"`
}
