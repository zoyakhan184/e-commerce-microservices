package repository

import (
	"auth-service/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	DB *gorm.DB
}

func NewPostgresUserRepo() *PostgresUserRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to auth DB: %v", err)
	}
	return &PostgresUserRepo{DB: db}
}

func (r *PostgresUserRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *PostgresUserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *PostgresUserRepo) UpdateUserPassword(userID, newHash string) error {
	return r.DB.Model(&models.User{}).Where("id = ?", userID).Update("password", newHash).Error
}

func (r *PostgresUserRepo) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("id = ?", userID).First(&user).Error
	return &user, err
}
