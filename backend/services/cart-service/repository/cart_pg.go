package repository

import (
	"cart-service/models"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresCartRepo struct {
	DB *gorm.DB
}

func NewPostgresCartRepo() *PostgresCartRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to DB")
	}

	db.AutoMigrate(&models.CartItem{})
	return &PostgresCartRepo{DB: db}
}

func (r *PostgresCartRepo) AddToCart(item *models.CartItem) error {
	item.AddedAt = time.Now()
	return r.DB.Save(item).Error
}

func (r *PostgresCartRepo) RemoveFromCart(userID, productID string) error {
	return r.DB.Delete(&models.CartItem{}, "user_id = ? AND product_id = ?", userID, productID).Error
}

func (r *PostgresCartRepo) UpdateCartItem(userID, productID string, quantity int) error {
	return r.DB.Model(&models.CartItem{}).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Update("quantity", quantity).Error
}

func (r *PostgresCartRepo) GetCart(userID string) ([]models.CartItem, error) {
	var items []models.CartItem
	err := r.DB.Where("user_id = ?", userID).Find(&items).Error
	return items, err
}

func (r *PostgresCartRepo) ClearCart(userID string) error {
	return r.DB.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error
}
