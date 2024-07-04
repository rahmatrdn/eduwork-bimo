package models

import (
	"database/sql"
	"eduwork-bimo/Final-Project-Simple-E-Commerce-System/entities"
)

type CartModel struct {
	db *sql.DB
}

func NewCartModel(db *sql.DB) *CartModel {
	return &CartModel{db: db}
}

func (m CartModel) CreateCart(userID int) (int, error) {
	cartModel := NewCartModel(m.db)
	userCart, err := cartModel.GetCartByUserID(userID)
	if userCart != nil {
		return userCart.CartID, nil
	}

	result, err := m.db.Exec("INSERT INTO cart (user_id) VALUES (?)", userID)
	if err != nil {
		return 0, err
	}

	cartID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(cartID), nil
}

func (m CartModel) GetCartByUserID(userID int) (*entities.Cart, error) {
	var cart entities.Cart
	row := m.db.QueryRow("SELECT cart_id, user_id FROM cart WHERE user_id = ?", userID)
	err := row.Scan(&cart.CartID, &cart.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No cart found
		}
		return nil, err
	}
	return &cart, nil
}

func (m CartModel) AddCartItem(cartID int, item *entities.CartItem) error {
	result, err := m.db.Exec("INSERT INTO cart_items (cart_id, product_id, quantity, price) VALUES (?, ?, ?, ?)", cartID, item.ProductID, item.Quantity, item.Price)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	item.CartItemID = int(id) // Assuming you have a CartItemID field in your CartItem struct
	return nil
}

func (m CartModel) GetCartItems(userID int) ([]entities.CartItem, error) {
	// fetch the cart_id for given user_id
	var cartID int
	err := m.db.QueryRow("SELECT cart_id FROM cart WHERE user_id = ?", userID).Scan(&cartID)
	if err != nil {
		return nil, err
	}

	// fetch items from the cart_items table using the cart_id
	rows, err := m.db.Query("SELECT cart_id, cart_item_id, product_id, quantity, price FROM cart_items WHERE cart_id = ?", cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cartItems := []entities.CartItem{}
	for rows.Next() {
		var item entities.CartItem
		err := rows.Scan(&item.CartID, &item.CartItemID, &item.ProductID, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, item)
	}

	return cartItems, nil
}

func (m CartModel) UpdateCartItem(cartID int, productID int, quantity int) error {
	_, err := m.db.Exec("UPDATE cart_items SET quantity = ? WHERE cart_id = ? AND product_id = ?", quantity, cartID, productID)
	return err
}

func (m CartModel) DeleteCartItem(cartID int, productID int) error {
	_, err := m.db.Exec("DELETE FROM cart_items WHERE cart_id = ? AND product_id = ?", cartID, productID)
	return err
}

func (m CartModel) ClearCartItems(cartID int) error {
	_, err := m.db.Exec("DELETE FROM cart_items WHERE cart_id = ?", cartID)
	return err
}
