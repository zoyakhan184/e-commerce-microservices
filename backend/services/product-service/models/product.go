package models

import "time"

type Product struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	CategoryId  string
	Price       float64
	Brand       string
	Quantity    int // <-- add this field
	CreatedAt   time.Time
	Images      []ProductImage `gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	ID         string `gorm:"primaryKey"`
	ProductID  string
	ImageURL   string
	FileName   string
	Base64Data string  `gorm:"-" json:"base64Data"` // ignore this field in DB
	Product    Product `gorm:"constraint:OnDelete:CASCADE;"`
}
