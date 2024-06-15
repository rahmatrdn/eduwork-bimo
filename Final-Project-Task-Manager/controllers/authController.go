package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"eduwork-bimo/Final-Project-Task-Manager/entities"
	"eduwork-bimo/Final-Project-Task-Manager/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entities.User
		json.NewDecoder(r.Body).Decode(&user)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)
		userModel := models.NewUserModel(db)
		err = userModel.Create(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entities.User
		json.NewDecoder(r.Body).Decode(&user)
		userModel := models.NewUserModel(db)
		dbUser, err := userModel.GetByUsername(user.Username)
		if err != nil {
			http.Error(w, "Invalid username for "+user.Username, http.StatusUnauthorized)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			UserID:   dbUser.ID,
			Username: dbUser.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		fmt.Printf("Debug: UserID in claims: %d\n", dbUser.ID) // Tambahkan debugging

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Hapus cookie 'token' dengan membuat cookie baru yang kadaluarsa
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: time.Now().Add(-time.Hour), // Waktu yang sudah lewat untuk membuat cookie tidak valid
		})

		// Atau, jika Anda menyimpan token di header Authorization, hapus header ini
		w.Header().Set("Authorization", "")

		fmt.Println(tokenString)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
	}
}
