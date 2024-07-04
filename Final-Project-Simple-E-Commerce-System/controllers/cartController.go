package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/entities"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/helpers"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/middlewares"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/models"

	"github.com/gorilla/mux"
)

func GetProductInCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId, err := middlewares.GetUserIdFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		cartModel := models.NewCartModel(db)
		cartItems, err := cartModel.GetCartItems(userId)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Cannot get your items in cart"})
			return
		}
		helper.ResponseJSON(w, http.StatusOK, cartItems)
	}
}

func AddProductToCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := middlewares.GetUserIdFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		cartModel := models.NewCartModel(db)
		userCartID, err := cartModel.CreateCart(userId)
		if err != nil {
			response := map[string]string{"error": "Cannot create cart"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		var request struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		}
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil || request.Quantity <= 0 {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}

		productModel := models.NewProductModel(db)
		userProduct, err := productModel.GetProductById(request.ProductID)
		if err != nil {
			response := map[string]string{"error": "Cannot get the product"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		var cartItem entities.CartItem
		cartItem.CartID = userCartID
		cartItem.Price = userProduct.Price
		cartItem.ProductID = userProduct.ProductID
		cartItem.Quantity = request.Quantity

		var curStock int
		curStock, err = productModel.GetProductStock(request.ProductID)
		if err != nil {
			response := map[string]string{"error": "Cannot get the current product stock"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		if curStock < cartItem.Quantity {
			response := map[string]string{"error": "Insufficient stock"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		err = cartModel.AddCartItem(userCartID, &cartItem)
		if err != nil {
			response := map[string]string{"error": "Cannot add the product to cart"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		cartItems, err := cartModel.GetCartItems(userId)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Cannot get your items in cart"})
			return
		}

		response := map[string]interface{}{
			"message":  "Product added successfully",
			"products": cartItems,
		}
		helper.ResponseJSON(w, http.StatusCreated, response)
	}
}

func UpdateProductInCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := middlewares.GetUserIdFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		cartModel := models.NewCartModel(db)
		userCart, err := cartModel.GetCartByUserID(userId)
		if err != nil {
			response := map[string]string{"error": "Cannot get your cart"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		productIDStr := mux.Vars(r)["id"]
		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
			return
		}

		var request struct {
			Quantity int `json:"quantity"`
		}
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil || request.Quantity <= 0 {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}

		var cartItem entities.CartItem
		cartItem.Quantity = request.Quantity

		err = cartModel.UpdateCartItem(userCart.CartID, productID, request.Quantity)
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Update failed"})
			return
		}

		cartItems, err := cartModel.GetCartItems(userId)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Cannot get your items in cart"})
			return
		}

		response := map[string]interface{}{
			"message":  "Product updated successfully",
			"products": cartItems,
		}
		helper.ResponseJSON(w, http.StatusCreated, response)
	}
}

func DeleteProductInCart(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := middlewares.GetUserIdFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		productIDStr := mux.Vars(r)["id"]
		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			response := map[string]string{"error": "Invalid product ID"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		cartModel := models.NewCartModel(db)
		userCart, err := cartModel.GetCartByUserID(userId)
		if err != nil {
			response := map[string]string{"error": "Cannot get your cart"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		err = cartModel.DeleteCartItem(userCart.CartID, productID)
		if err != nil {
			response := map[string]string{"error": "failed to delete product"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		helper.ResponseJSON(w, http.StatusNoContent, nil)
	}
}
