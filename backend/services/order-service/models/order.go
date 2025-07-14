package models

import "time"

type Order struct {
	ID            string    `gorm:"primaryKey"`
	UserID        string
	TotalAmount   float64
	OrderStatus   string
	PaymentStatus string
	CreatedAt     time.Time
	OrderItems    []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        string `gorm:"primaryKey"`
	OrderID   string
	ProductID string
	Quantity  int
	Size      string
	Color     string
	Price     float64
}
