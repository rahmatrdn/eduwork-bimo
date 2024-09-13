package routes

import (
	"database/sql"

	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/controllers"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.Register(db)).Methods("POST")
	router.HandleFunc("/login", controllers.Login(db)).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout()).Methods("POST")
	router.HandleFunc("/reset-password", controllers.ResetPassword(db)).Methods("PUT")
	router.HandleFunc("/products", controllers.GetProducts(db)).Methods("GET")
	router.HandleFunc("/products", controllers.CreateProduct(db)).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)

	api.HandleFunc("/profile", controllers.GetProfile(db)).Methods("GET")
	api.HandleFunc("/profile", controllers.UpdateProfile(db)).Methods("PUT")
	api.HandleFunc("/profile", controllers.DeleteUser(db)).Methods("DELETE")

	api.HandleFunc("/cart", controllers.GetProductInCart(db)).Methods("GET")
	api.HandleFunc("/cart", controllers.AddProductToCart(db)).Methods("POST")
	api.HandleFunc("/cart/{id}", controllers.UpdateProductInCart(db)).Methods("PUT")
	api.HandleFunc("/cart/{id}", controllers.DeleteProductInCart(db)).Methods("DELETE")

	api.HandleFunc("/order", controllers.CreateOrder(db)).Methods("POST")
	api.HandleFunc("/order", controllers.GetOrderSummary(db)).Methods("GET")
	api.HandleFunc("/order", controllers.UpdateOrderStatus(db)).Methods("PUT")
	api.HandleFunc("/order/history", controllers.GetOrderHistory(db)).Methods("GET")

	return router
}
