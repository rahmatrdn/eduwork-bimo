package entities

type CartItem struct {
	CartItemID int     `json:"cart_item_id"`
	CartID     int     `json:"cart_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}
