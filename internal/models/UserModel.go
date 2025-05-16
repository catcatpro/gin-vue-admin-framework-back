package models

import (
	"fmt"
	"gin_vue_admin_framework/common"
	"gin_vue_admin_framework/utils"
)

type User struct {
	*Basic
	Username string `json:"username" gorm:"NOT NULL"`
	Password string `json:"password" gorm:"NOT NULL"`
}

// func (User) TableName() string {
// 	return "user"
// }

// 用户登录
func (u *User) UserLogin() error {
	//var user User
	fmt.Println(u.Username, u.Password)
	u.Password, _ = utils.HashPassword(u.Password)
	res := common.COM_DB.Where("username = ?", u.Username).Where("password", u.Password).First(&u)
	// 临时登录逻辑

	return res.Error
}

// 新建用户
func (u *User) CreateUser() (uint, error) {
	var saveData User
	saveData = *u
	saveData.Password, _ = utils.HashPassword(saveData.Password)
	res := common.COM_DB.Create(&saveData)
	return saveData.ID, res.Error
}
