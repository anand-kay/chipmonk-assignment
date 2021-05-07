package libs

import (
	"errors"
)

func ValidateSessionToken(sessionToken string) error {
	if sessionToken == "" {
		return errors.New("Session token can't be empty")
	}

	if len(sessionToken) != 10 {
		return errors.New("Invalid session token")
	}

	for _, chr := range sessionToken {
		if !(chr >= 97 && chr <= 122) && !(chr >= 48 && chr <= 57) {
			return errors.New("Invalid session token")
		}
	}

	return nil
}
