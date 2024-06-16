package helper

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func EncryptPass(userPass string) (encryptedPass string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	encryptedPass = string(hashedPassword)

	return encryptedPass, nil
}
