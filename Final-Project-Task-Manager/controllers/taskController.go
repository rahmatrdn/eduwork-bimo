package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"eduwork-bimo/Final-Project-Task-Manager/entities"
	"eduwork-bimo/Final-Project-Task-Manager/helper"
	"eduwork-bimo/Final-Project-Task-Manager/middlewares"
	"eduwork-bimo/Final-Project-Task-Manager/models"

	"github.com/gorilla/mux"
)

func GetTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middlewares.GetUserIDFromToken(r)
		if err != nil {
			response := map[string]string{"error": "Unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		taskModel := models.NewTaskModel(db)
		tasks, err := taskModel.GetAll(userID)
		if err != nil {
			response := map[string]string{"error": "Failed to fetch tasks"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		helper.ResponseJSON(w, http.StatusOK, tasks)
	}
}

func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middlewares.GetUserIDFromToken(r)
		if err != nil {
			response := map[string]string{"error": "Unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		var task entities.Task
		err = json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			response := map[string]string{"error": "Invalid request payload"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		// Contoh validasi sederhana untuk payload
		if task.Title == "" {
			response := map[string]string{"error": "Title is required"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		task.UserID = userID
		taskModel := models.NewTaskModel(db)
		err = taskModel.Create(&task)
		if err != nil {
			response := map[string]string{"error": "Failed to create task"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := map[string]interface{}{
			"message": "Task created successfully",
			"task":    task,
		}
		helper.ResponseJSON(w, http.StatusCreated, response)
	}
}

func UpdateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middlewares.GetUserIDFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		taskIDStr := mux.Vars(r)["id"]
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
			return
		}

		var task entities.Task
		err = json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}

		task.ID = taskID
		task.UserID = userID

		taskModel := models.NewTaskModel(db)
		err = taskModel.Update(&task)
		if err != nil {
			helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		response := map[string]interface{}{
			"message": "Task updated successfully",
			"task":    task,
		}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}

func DeleteTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middlewares.GetUserIDFromToken(r)
		if err != nil {
			response := map[string]string{"error": "failed to get user ID"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		taskIDStr := mux.Vars(r)["id"]
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			response := map[string]string{"error": "Invalid task ID"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		taskModel := models.NewTaskModel(db)
		err = taskModel.Delete(taskID, userID)
		if err != nil {
			response := map[string]string{"error": "failed to delete task"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		helper.ResponseJSON(w, http.StatusNoContent, nil)
	}
}
