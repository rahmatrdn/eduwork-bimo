package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/entities"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/helpers"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/middlewares"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/models"

	"github.com/gorilla/mux"
)

func GetProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		productModel := models.NewProductModel(db)
		products, err := productModel.GetAll()
		if err != nil {
			response := map[string]string{"error": "Failed to fetch products"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		helper.ResponseJSON(w, http.StatusOK, products)
	}
}

func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, err := middlewares.GetUserRoleFromToken(r)
		fmt.Println(role)
		if err != nil {
			response := map[string]string{"error": "Unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		if role != "admin" {
			response := map[string]string{"error": "Forbidden"}
			helper.ResponseJSON(w, http.StatusForbidden, response)
			return
		}

		var product entities.Product
		err = json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			response := map[string]string{"error": "Invalid request payload"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		// Contoh validasi sederhana untuk payload
		if product.ProductName == "" {
			response := map[string]string{"error": "Product name is required"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		productModel := models.NewProductModel(db)
		err = productModel.Create(&product)
		if err != nil {
			response := map[string]string{"error": "Failed to create product"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := map[string]interface{}{
			"message": "Product created successfully",
			"product": product,
		}
		helper.ResponseJSON(w, http.StatusCreated, response)
	}
}

func UpdateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, err := middlewares.GetUserRoleFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		if role != "admin" {
			response := map[string]string{"error": "Forbidden"}
			helper.ResponseJSON(w, http.StatusForbidden, response)
			return
		}

		productIDStr := mux.Vars(r)["id"]
		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
			return
		}

		var product entities.Product
		err = json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}

		product.ProductID = productID

		productModel := models.NewProductModel(db)
		err = productModel.Update(&product)
		if err != nil {
			helper.ResponseJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		response := map[string]interface{}{
			"message": "Product updated successfully",
			"product": product,
		}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}

func DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, err := middlewares.GetUserRoleFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		if role != "admin" {
			response := map[string]string{"error": "Forbidden"}
			helper.ResponseJSON(w, http.StatusForbidden, response)
			return
		}

		productIDStr := mux.Vars(r)["id"]
		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			response := map[string]string{"error": "Invalid product ID"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		productModel := models.NewProductModel(db)
		err = productModel.Delete(productID)
		if err != nil {
			response := map[string]string{"error": "failed to delete product"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		helper.ResponseJSON(w, http.StatusNoContent, nil)
	}
}
