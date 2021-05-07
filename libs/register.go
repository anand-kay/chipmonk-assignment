package libs

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func ValidateForm(username string, password string) error {
	if !validateName(username) {
		return errors.New("Invalid username")
	}

	if !validatePassword(password) {
		return errors.New("Invalid password")
	}

	return nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func validateName(name string) bool {
	rxp := regexp.MustCompile("^[A-Za-z0-9]+$")

	return rxp.MatchString(name)
}

func validatePassword(password string) bool {
	rxp := regexp.MustCompile("^[A-Za-z0-9_@.# &+-]{8,16}$")

	return rxp.MatchString(password)
}
