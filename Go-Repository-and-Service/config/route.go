package config

import (
	"eduwork-bimo/Go-Repository-and-Service/modules/product"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Add modules route
	product.ProductRoutes(app, db)

	port := os.Getenv("PORT")
	portStr := ":" + port
	app.Listen(portStr)
}
