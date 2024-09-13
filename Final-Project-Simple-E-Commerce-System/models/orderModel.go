package models

import (
	"database/sql"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/entities"
)

type OrderModel struct {
	db *sql.DB
}

func NewOrderModel(db *sql.DB) *OrderModel {
	return &OrderModel{db: db}
}

func (m OrderModel) CreateOrder(userID int, items []entities.OrderItem) (int, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return 0, err
	}

	var totalAmount float64
	for _, item := range items {
		totalAmount += item.Price * float64(item.Quantity)
	}

	result, err := tx.Exec("INSERT INTO orders (user_id, total_amount, status) VALUES (?, ?, 'pending')", userID, totalAmount)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, item := range items {
		item.OrderID = int(orderID)
		_, err := tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)", item.OrderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(orderID), nil
}

func (m OrderModel) CreateOrderFromCart(userID int) (int, error) {
	// Retrieve cart items for the user
	cartModel := NewCartModel(m.db)
	cartItems, err := cartModel.GetCartItems(userID)
	if err != nil {
		return 0, err
	}

	// Convert cart items to order items
	orderItems := []entities.OrderItem{}
	for _, item := range cartItems {
		orderItem := entities.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		orderItems = append(orderItems, orderItem)
	}

	// Create order using the cart items
	orderID, err := m.CreateOrder(userID, orderItems)
	if err != nil {
		return 0, err
	}

	userCart, err := cartModel.GetCartByUserID(userID)
	if err != nil {
		return 0, err
	}

	// Clear cart items after creating order
	err = cartModel.ClearCartItems(userCart.CartID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func (m OrderModel) GetOrderSummary(userID int, status string) (*entities.Order, error) {
	var orderSum entities.Order
	row := m.db.QueryRow("SELECT order_id, user_id, total_amount, status FROM orders WHERE user_id = ? AND status = ?", userID, status)
	err := row.Scan(&orderSum.OrderID, &orderSum.UserID, &orderSum.TotalAmount, &orderSum.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No order found
		}
		return nil, err
	}
	return &orderSum, nil
}

func (m OrderModel) GetOrderIdFromUserId(userID int, status string) (int, error) {
	var orderID int
	err := m.db.QueryRow("SELECT order_id FROM orders WHERE user_id = ? AND status = ?", userID, status).Scan(&orderID)
	if err != nil {
		return -1, err
	}

	return orderID, nil
}

func (m OrderModel) GetOrderIdsFromUserId(userID int, status string) ([]int, error) {
	rows, err := m.db.Query("SELECT order_id FROM orders WHERE user_id = ? AND status = ?", userID, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderIDs []int
	for rows.Next() {
		var orderID int
		if err := rows.Scan(&orderID); err != nil {
			return nil, err
		}
		orderIDs = append(orderIDs, orderID)
	}

	// Check if any rows were returned
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(orderIDs) == 0 {
		return nil, sql.ErrNoRows
	}

	return orderIDs, nil
}

func (m OrderModel) GetOrderItem(orderID int) ([]entities.OrderItem, error) {
	// fetch the items from the order_items table using the order_id
	rows, err := m.db.Query("SELECT order_item_id, order_id, product_id, quantity, price FROM order_items WHERE order_id = ?", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderItems := []entities.OrderItem{}
	for rows.Next() {
		var item entities.OrderItem
		err := rows.Scan(&item.OrderItemID, &item.OrderID, &item.ProductID, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, item)
	}

	return orderItems, nil
}

func (m OrderModel) GetOrderStatus(orderID int) (string, error) {
	var status string
	err := m.db.QueryRow("SELECT status FROM orders WHERE order_id = ?", orderID).Scan(&status)
	if err != nil {
		return "", err
	}

	return status, nil
}

func (m OrderModel) UpdateOrderStatus(orderID int, status string) error {
	_, err := m.db.Exec("UPDATE orders SET status = ? WHERE order_id = ?", status, orderID)
	return err
}

func (m OrderModel) DecreaseProductStock(orderItem *entities.OrderItem) error {
	_, err := m.db.Exec("UPDATE products SET stock = stock - ? WHERE product_id = ?", orderItem.Quantity, orderItem.ProductID)
	return err
}
