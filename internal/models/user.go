package models

type User struct {
	Baseic
	Username string `json:"username"`
	Password string `json:"password"`
}
