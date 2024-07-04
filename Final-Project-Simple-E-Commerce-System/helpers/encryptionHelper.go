package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPass(userPass string) (encryptedPass string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	encryptedPass = string(hashedPassword)

	return encryptedPass, nil
}
