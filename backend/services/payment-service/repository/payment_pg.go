package repository

import (
	"log"
	"os"
	"payment-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PaymentRepo struct {
	DB *gorm.DB
}

func NewPaymentRepo() *PaymentRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Payment{}); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	return &PaymentRepo{DB: db}
}

func (r *PaymentRepo) Create(payment *models.Payment) error {
	return r.DB.Create(payment).Error
}

func (r *PaymentRepo) UpdateStatus(id, status, txnRef string) error {
	return r.DB.Model(&models.Payment{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":  status,
			"txn_ref": txnRef,
			"paid_at": gorm.Expr("NOW()"),
		}).Error
}

func (r *PaymentRepo) GetByID(id string) (*models.Payment, error) {
	var p models.Payment
	err := r.DB.Where("id = ?", id).First(&p).Error
	return &p, err
}

func (r *PaymentRepo) GetByOrderID(orderID string) (*models.Payment, error) {
	var p models.Payment
	err := r.DB.Where("order_id = ?", orderID).First(&p).Error
	return &p, err
}
