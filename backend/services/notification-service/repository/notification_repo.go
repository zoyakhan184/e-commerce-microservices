package repository

import (
	"log"
	"notification-service/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	DB *gorm.DB
}

func NewNotificationRepo() *NotificationRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Failed:", err)
	}

	db.AutoMigrate(&models.Notification{})
	return &NotificationRepo{DB: db}
}

func (r *NotificationRepo) Save(n *models.Notification) error {
	return r.DB.Create(n).Error
}
