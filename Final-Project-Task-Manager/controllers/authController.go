package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"eduwork-bimo/Final-Project-Task-Manager/entities"
	"eduwork-bimo/Final-Project-Task-Manager/helper"
	"eduwork-bimo/Final-Project-Task-Manager/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("ashdjqy9283409bsdklkg8hda02")

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entities.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response := map[string]string{"error": "Invalid request payload"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			response := map[string]string{"error": "Failed to hash password"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		user.Password = string(hashedPassword)
		userModel := models.NewUserModel(db)
		err = userModel.Create(&user)
		if err != nil {
			response := map[string]string{"error": "Failed to create user"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := map[string]string{"message": "User registered successfully"}
		helper.ResponseJSON(w, http.StatusCreated, response)
	}
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// validasi user
		var user entities.User
		json.NewDecoder(r.Body).Decode(&user)
		userModel := models.NewUserModel(db)
		dbUser, err := userModel.GetByUsername(user.Username)
		if err != nil {
			response := map[string]string{"error": "Invalid username for " + user.Username}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		defer r.Body.Close()

		// validasi password
		err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
		if err != nil {
			response := map[string]string{"error": "Invalid password"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// pembuatan token
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			UserID:   dbUser.ID,
			Username: dbUser.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			response := map[string]string{"error": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    tokenString,
			HttpOnly: true,
		})

		response := map[string]string{"message": "Login success"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// hapus token yang ada di cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    "",
			HttpOnly: true,
			MaxAge:   -1,
		})

		response := map[string]string{"message": "Logout success"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}
