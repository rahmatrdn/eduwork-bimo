package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRoutes(app *fiber.App, db *gorm.DB) {
	productRepository := NewProductRepository(db)
	productService := NewProductService(productRepository)
	productHandler := NewProductHandler(productService)

	app.Get("/api/products", productHandler.GetAllProducts)
	app.Post("/api/products", productHandler.CreateProduct)
	app.Get("/api/products/:id", productHandler.GetProductById)
	app.Put("/api/products/:id", productHandler.UpdateProduct)
	app.Delete("/api/products/:id", productHandler.DeleteProduct)
}
