package routes

import (
	"database/sql"

	"eduwork-bimo/Final-Project-User-Manager/controllers"
	"eduwork-bimo/Final-Project-User-Manager/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.Register(db)).Methods("POST")
	router.HandleFunc("/login", controllers.Login(db)).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout()).Methods("POST")
	router.HandleFunc("/reset-password", controllers.ResetPassword(db)).Methods("PUT")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)
	api.HandleFunc("/profile", controllers.GetProfile(db)).Methods("GET")
	api.HandleFunc("/profile", controllers.UpdateProfile(db)).Methods("PUT")
	api.HandleFunc("/profile", controllers.DeleteUser(db)).Methods("DELETE")

	return router
}
