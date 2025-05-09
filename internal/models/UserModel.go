package models

import (
	"fmt"
	"gin_vue_admin_framework/common"
)

type User struct {
	*Basic
	Username string `json:"username"`
	Password string `json:"password"`
}

// func (User) TableName() string {
// 	return "user"
// }

// 用户登录
func (u *User) UserLogin() bool {
	var res User
	fmt.Println(u.Username, u.Password)

	common.COM_DB.Where("username = ?", u.Username).Where("password", u.Password).First(&res)
	// 临时登录逻辑
	fmt.Println("username " + res.Username + " password " + res.Password)

	return res.Username != ""
}
