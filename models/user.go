package models

import "github.com/anand-kay/chipmonk-assignment/db"

func (user *User) RegisterUser() error {
	db.Users[user.Username] = user.Password

	return nil
}

func (user *User) CheckUserExists() bool {
	if _, ok := db.Users[user.Username]; ok {
		return true
	}

	return false
}

func (user *User) RetreiveStoredHash() string {
	return db.Users[user.Username]
}
