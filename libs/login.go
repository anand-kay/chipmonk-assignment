package libs

import "golang.org/x/crypto/bcrypt"

func CheckHash(inputPwd string, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(inputPwd))
	if err != nil {
		return false
	}

	return true
}
