package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID                int       `json:"id`
	Nama_Pengguna     string    `json:"nama_pengguna`
	Email_Pengguna    string    `json:"email_pengguna`
	Password          string    `json:"password`
	Tanggal_Pembuatan time.Time `json:"tanggal_pembuatan`
}

type Product struct {
	ID                int       `json:"id`
	Nama_Produk       string    `json:"nama_produk`
	Deskripsi         string    `json:"deskripsi`
	Harga             float64   `json:"harga`
	Stok              int       `json:"stok`
	Tanggal_Pembuatan time.Time `json:"tanggal_pembuatan`
}

// User input when login
type LoginRequest struct {
	Nama_Pengguna string `json:"nama_pengguna`
	Password      string `json:"password"`
}

// Response after user logged-in
type LoginResponse struct {
	Token string `json:"token"`
}

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/eduwork?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	app := fiber.New()

	// Create a user
	app.Post("/api/users/register", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		// Check if the username already exists
		var existingUser User
		if err := db.Where("nama_pengguna = ?", user.Nama_Pengguna).First(&existingUser).Error; err == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Username already exists"})
		}

		// Check if the email already exists
		if err := db.Where("email_pengguna = ?", user.Email_Pengguna).First(&existingUser).Error; err == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Email already exists"})
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)

		db.Create(&user)
		return c.JSON(user)
	})

	// User authentication
	app.Post("/api/users/login", func(c *fiber.Ctx) error {
		var loginReq LoginRequest
		if err := c.BodyParser(&loginReq); err != nil {
			return err
		}

		var user User
		if err := db.Where("nama_pengguna = ?", loginReq.Nama_Pengguna).First(&user).Error; err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password.")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password.")
		}

		token := "here's your token"

		return c.JSON(LoginResponse{Token: token})
	})

	// Get all products
	app.Get("/api/products", func(c *fiber.Ctx) error {
		var products []Product
		if err := db.Find(&products).Error; err != nil {
			return err
		}
		return c.JSON(products)
	})

	// Create a product
	app.Post("/api/products", func(c *fiber.Ctx) error {
		var product Product
		if err := c.BodyParser(&product); err != nil {
			return err
		}
		db.Create(&product)
		return c.JSON(product)
	})

	// Get a single product
	app.Get("/api/products/:id", func(c *fiber.Ctx) error {
		var product Product
		id := c.Params("id")
		if err := db.First(&product, id).Error; err != nil {
			return err
		}
		return c.JSON(product)
	})

	// Update a product
	app.Put("/api/products/:id", func(c *fiber.Ctx) error {
		var product Product
		id := c.Params("id")
		if err := db.First(&product, id).Error; err != nil {
			return err
		}
		if err := c.BodyParser(&product); err != nil {
			return err
		}

		db.Save(&product)
		return c.JSON(product)
	})

	// Delete a product
	app.Delete("/api/products/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := db.Delete(&Product{}, id).Error; err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

	port := "3000"
	app.Listen(":" + port)
}
