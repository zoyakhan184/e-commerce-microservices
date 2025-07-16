package repository

import (
	"order-service/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo() *OrderRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Order{}, &models.OrderItem{})
	return &OrderRepo{DB: db}
}

func (r *OrderRepo) CreateOrder(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepo) GetOrdersByUser(userID string) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Where("user_id = ?", userID).Preload("OrderItems").Find(&orders).Error
	return orders, err
}

func (r *OrderRepo) GetOrderByID(orderID string) (models.Order, error) {
	var order models.Order
	err := r.DB.Preload("OrderItems").First(&order, "id = ?", orderID).Error
	return order, err
}

func (r *OrderRepo) UpdateOrderStatus(orderID, status string) error {
	return r.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("order_status", status).Error
}

func (r *OrderRepo) UpdatePaymentStatus(orderID, status string) error {
	return r.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("payment_status", status).Error
}

func (r *OrderRepo) GetOrderWithItems(orderID string) (models.Order, error) {
	var order models.Order
	err := r.DB.Preload("OrderItems").Where("id = ?", orderID).First(&order).Error
	return order, err
}

func (r *OrderRepo) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("OrderItems").Order("created_at desc").Find(&orders).Error
	return orders, err
}
