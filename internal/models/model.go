package models

import "gorm.io/gorm"

type Basic struct {
	*gorm.Model
	ID uint `json:"id" gorm:"primarykey"`
}
