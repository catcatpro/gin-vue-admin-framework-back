package core

import "time"

type User struct {
	id       int
	username string
	password string
	token    string
	createAt time.Time
	updateAt time.Time
}

type UserInterface interface {
	getUserById(id int) User
	getByUsername(username string) User
	checkPassword(username string, password string) (User, error)
	checkToken(token string) User
}
