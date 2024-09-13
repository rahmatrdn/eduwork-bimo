package main

import (
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/config"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := config.DBConn()
	defer db.Close()

	router := routes.SetupRoutes(db)

	fmt.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
