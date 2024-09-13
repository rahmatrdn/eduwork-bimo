package config

import (
	"github.com/dgrijalva/jwt-go"
)

var JWT_KEY = []byte("ashdjqy9283409bsdklkg8hda03")

type JWTClaim struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
