package models

type User struct {
	Basic
	Username string `json:"username"`
	Password string `json:"password"`
}
