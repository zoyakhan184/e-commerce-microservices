package models

type Order struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	TotalAmount   float64 `json:"total_amount"`
	OrderStatus   string  `json:"order_status"`
	PaymentStatus string  `json:"payment_status"`
	CreatedAt     string  `json:"created_at"`
}
