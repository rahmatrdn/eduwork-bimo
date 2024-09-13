package main

import (
	"eduwork-bimo/Go-Repository-and-Service/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()

	// Connect to database
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	// Setup routes
	config.SetupRoutes(app, db)
}
