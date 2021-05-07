package models

import (
	"math/rand"
	"time"

	"github.com/anand-kay/chipmonk-assignment/db"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func (user *User) CreateNewSession() string {
	rand.Seed(time.Now().UnixNano())

	sessionToken := genSessionToken(10)

	db.Sessions[user.Username] = sessionToken

	return sessionToken
}

func IsUserActive(inputToken string) bool {
	for _, storedToken := range db.Sessions {
		if inputToken == storedToken {
			return true
		}
	}

	return false
}

func FetchActiveUsers() []string {
	var activeUsers []string

	for username := range db.Sessions {
		activeUsers = append(activeUsers, username)
	}

	return activeUsers
}

func genSessionToken(n int) string {
	b := make([]rune, n)
	lenRunes := len(letterRunes)
	for i := range b {
		b[i] = letterRunes[rand.Intn(lenRunes)]
	}
	return string(b)
}
