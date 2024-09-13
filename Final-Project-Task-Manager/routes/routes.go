package routes

import (
	"database/sql"

	"eduwork-bimo/Final-Project-Task-Manager/controllers"
	"eduwork-bimo/Final-Project-Task-Manager/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.Register(db)).Methods("POST")
	router.HandleFunc("/login", controllers.Login(db)).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout()).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)
	api.HandleFunc("/tasks", controllers.GetTasks(db)).Methods("GET")
	api.HandleFunc("/tasks", controllers.CreateTask(db)).Methods("POST")
	api.HandleFunc("/tasks/{id}", controllers.UpdateTask(db)).Methods("PUT")
	api.HandleFunc("/tasks/{id}", controllers.DeleteTask(db)).Methods("DELETE")

	return router
}
