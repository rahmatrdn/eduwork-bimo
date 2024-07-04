package entities

type Order struct {
	OrderID     int     `json:"order_id"`
	UserID      int     `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}
