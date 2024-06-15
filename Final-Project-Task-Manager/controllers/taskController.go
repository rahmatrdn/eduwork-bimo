package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"eduwork-bimo/Final-Project-Task-Manager/entities"
	"eduwork-bimo/Final-Project-Task-Manager/middlewares"
	"eduwork-bimo/Final-Project-Task-Manager/models"

	"github.com/gorilla/mux"
)

func GetTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middlewares.GetUserIDFromToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		taskModel := models.NewTaskModel(db)
		tasks, err := taskModel.GetAll(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tasks)
	}
}

func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middlewares.GetUserIDFromToken(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var task entities.Task
		err = json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		task.UserID = userID
		taskModel := models.NewTaskModel(db)
		err = taskModel.Create(&task)
		if err != nil {
			fmt.Println(userID)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

func UpdateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middlewares.GetUserIDFromToken(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		taskIDStr := mux.Vars(r)["id"]
		taskID, err := strconv.Atoi(taskIDStr)
		fmt.Printf("Debug: TaskID : %s\n", taskIDStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		var task entities.Task
		err = json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		task.ID = taskID
		task.UserID = userID

		taskModel := models.NewTaskModel(db)
		err = taskModel.Update(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := middlewares.GetUserIDFromToken(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		taskIDStr := mux.Vars(r)["id"]
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		taskModel := models.NewTaskModel(db)
		err = taskModel.Delete(taskID, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
