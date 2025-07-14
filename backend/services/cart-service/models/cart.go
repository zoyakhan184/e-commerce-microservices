package models

import "time"

type CartItem struct {
    UserID    string `gorm:"primaryKey"`
    ProductID string `gorm:"primaryKey"`
    Size      string
    Color     string
    Quantity  int
    AddedAt   time.Time
}
