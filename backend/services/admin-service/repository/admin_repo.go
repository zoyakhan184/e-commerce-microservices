package repository

import (
	"log"
	"os"
	"sync"
	"admin-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

var (
	activityStore []models.ActivityLog
	storeMutex    sync.Mutex // for safe concurrent access
)

func NewAdminRepo() *AdminRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connect failed: %v", err)
	}
	return &AdminRepo{DB: db}
}

func (r *AdminRepo) CountUsers() (int64, error) {
	var count int64
	err := r.DB.Table("users").Count(&count).Error
	return count, err
}

func (r *AdminRepo) CountOrders() (int64, error) {
	var count int64
	err := r.DB.Table("orders").Count(&count).Error
	return count, err
}

func (r *AdminRepo) SumRevenue() (float64, error) {
	var sum float64
	err := r.DB.Table("orders").Select("SUM(total_amount)").Row().Scan(&sum)
	return sum, err
}

func (r *AdminRepo) ListUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Table("users").Select("id as user_id, name, email").Scan(&users).Error
	return users, err
}

func (r *AdminRepo) ListOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Table("orders").Scan(&orders).Error
	return orders, err
}

func (r *AdminRepo) ListLowStockThreshold(threshold int) ([]models.Inventory, error) {
	var items []models.Inventory
	err := r.DB.Table("inventory").Where("quantity <= ?", threshold).Scan(&items).Error
	return items, err
}

func (r *AdminRepo) ListRecentUsers(limit int) ([]models.User, error) {
	var users []models.User
	err := r.DB.Raw(`SELECT id as user_id, name, email, created_at FROM users ORDER BY created_at DESC LIMIT ?`, limit).Scan(&users).Error
	return users, err
}

func (r *AdminRepo) ListRecentOrders(limit int) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Raw(`SELECT id, user_id, total_amount, order_status, payment_status, created_at FROM orders ORDER BY created_at DESC LIMIT ?`, limit).Scan(&orders).Error
	return orders, err
}

func (r *AdminRepo) AddActivity(log models.ActivityLog) {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	activityStore = append([]models.ActivityLog{log}, activityStore...)
	if len(activityStore) > 10 {
		activityStore = activityStore[:10]
	}
}

func (r *AdminRepo) GetRecentActivity() []models.ActivityLog {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	// Return a shallow copy to avoid race conditions
	clone := make([]models.ActivityLog, len(activityStore))
	copy(clone, activityStore)
	return clone
}

