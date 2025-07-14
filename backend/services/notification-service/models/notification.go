package models

import "time"

type Notification struct {
	ID        string    `gorm:"primaryKey"`
	UserID    string
	Title     string
	Message   string
	Read      bool
	CreatedAt time.Time
}
