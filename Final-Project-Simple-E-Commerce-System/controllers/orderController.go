package controllers

import (
	"database/sql"
	"net/http"

	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/helpers"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/middlewares"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/models"
)

func GetOrderSummary(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId, err := middlewares.GetUserIdFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		orderModel := models.NewOrderModel(db)

		orderID, err := orderModel.GetOrderIdFromUserId(userId, "pending")
		if err != nil {
			response := map[string]string{"error": "Cannot find your order"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		orderItems, err := orderModel.GetOrderItem(orderID)
		if err != nil {
			response := map[string]string{"error": "Failed to fetch your order"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		orderSum, err := orderModel.GetOrderSummary(userId, "pending")
		if err != nil {
			response := map[string]string{"error": "Failed to count your total"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := map[string]interface{}{
			"message": "Order Summary",
			"order":   orderItems,
			"total":   orderSum.TotalAmount,
			"status":  orderSum.Status,
		}
		helper.ResponseJSON(w, http.StatusCreated, response)
	}
}

func CreateOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId, err := middlewares.GetUserIdFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		orderModel := models.NewOrderModel(db)

		orderID, err := orderModel.CreateOrderFromCart(userId)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Cannot make your order"})
			return
		}

		orderItems, err := orderModel.GetOrderItem(orderID)
		if err != nil {
			response := map[string]string{"error": "Failed to fetch your order"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := map[string]interface{}{
			"message": "Order created successfully",
			"order":   orderItems,
		}
		helper.ResponseJSON(w, http.StatusCreated, response)
	}
}

func UpdateOrderStatus(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := middlewares.GetUserIdFromToken(r)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			return
		}

		orderModel := models.NewOrderModel(db)
		orderID, err := orderModel.GetOrderIdFromUserId(userId, "pending")
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Cannot fetch your order ID"})
			return
		}

		orderItems, err := orderModel.GetOrderItem(orderID)
		if err != nil {
			helper.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"error": "Failed to fetch your order"})
			return
		}

		productModel := models.NewProductModel(db)
		var insufficientStock bool

		for _, item := range orderItems {
			curStock, err := productModel.GetProductStock(item.ProductID)
			if err != nil {
				response := map[string]string{"error": "Cannot get the current product stock"}
				helper.ResponseJSON(w, http.StatusBadRequest, response)
				return
			}

			if curStock < item.Quantity {
				insufficientStock = true
				break
			}

			err = orderModel.DecreaseProductStock(&item)
			if err != nil {
				response := map[string]string{"error": "Failed to update product stock"}
				helper.ResponseJSON(w, http.StatusBadRequest, response)
				return
			}
		}

		if insufficientStock {
			response := map[string]string{"error": "Insufficient stock for one or more products"}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}

		err = orderModel.UpdateOrderStatus(orderID, "completed")
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Update failed"})
			return
		}

		status, err := orderModel.GetOrderStatus(orderID)
		if err != nil {
			helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch order status"})
			return
		}

		response := map[string]interface{}{
			"message": "Order updated successfully",
			"order":   orderItems,
			"status":  status,
		}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}
