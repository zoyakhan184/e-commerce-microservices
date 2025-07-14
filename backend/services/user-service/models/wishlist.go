package models

type Wishlist struct {
	UserID    string `gorm:"primaryKey;column:user_id"`
	ProductID string `gorm:"primaryKey;column:product_id"`
}
