package models

import "time"

type Review struct {
	ID        string    `gorm:"primaryKey"`
	UserID    string    `gorm:"not null"`
	ProductID string    `gorm:"not null"`
	Rating    int       `gorm:"not null"`
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
