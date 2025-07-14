package models

import "time"

type Payment struct {
	ID        string `gorm:"primaryKey"`
	OrderID   string
	UserID    string
	Gateway   string
	Status    string
	TxnRef    string
	PaidAt    *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
