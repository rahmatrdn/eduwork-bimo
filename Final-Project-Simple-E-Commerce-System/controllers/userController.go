package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/entities"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/helpers"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/middlewares"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/models"
)

func GetProfile(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := middlewares.GetUsernameFromToken(r)
		if err != nil {
			response := map[string]string{"error": "Unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		userModel := models.NewUserModel(db)
		profile, err := userModel.GetByUsername(username)
		if err != nil {
			response := map[string]string{"error": "Failed to fetch profile"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		helper.ResponseJSON(w, http.StatusOK, profile)
	}
}

func UpdateProfile(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := middlewares.GetUserIdFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		var user entities.User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}

		user.UserID = userId

		userModel := models.NewUserModel(db)
		err = userModel.Update(&user)
		if err != nil {
			helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		response := map[string]interface{}{
			"message": "Profile updated successfully",
			"profile": user,
		}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middlewares.GetUserIdFromToken(r)
		if err != nil {
			response := map[string]string{"error": "failed to get user ID"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		userModel := models.NewUserModel(db)
		err = userModel.DeleteUser(userID)
		if err != nil {
			response := map[string]string{"error": "failed to delete task"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		// Hapus token dari cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    "",
			HttpOnly: true,
			MaxAge:   -1,
		})

		response := map[string]string{"message": "User deleted and logged out"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}

func ResetPassword(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user entities.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}

		userModel := models.NewUserModel(db)
		err = userModel.ResetPassword(&user)
		if err != nil {
			helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		response := map[string]interface{}{
			"message": "Profile updated successfully",
			"profile": user,
		}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}
