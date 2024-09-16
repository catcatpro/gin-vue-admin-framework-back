package models

type User struct {
	Basic
	Username string `json:"username"`
	Password string `json:"password"`
}

/**
*用户登录
**/
func (u *User) UserLogin() bool {
	//TODO 登录逻辑 -- 查询数据库

	//TODO 临时登录逻辑
	if u.Username != "admin" && u.Password != "admin" {
		return false
	}
	return true
}
