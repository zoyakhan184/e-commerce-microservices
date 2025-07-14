package models

import "time"

type Inventory struct {
	SkuID     string `gorm:"primaryKey"`
	ProductID string
	Size      string
	Color     string
	Quantity  int
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
