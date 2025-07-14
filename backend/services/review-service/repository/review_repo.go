package repository

import (
	"os"
	"review-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ReviewRepo struct {
	DB *gorm.DB
}

func NewReviewRepo() *ReviewRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to DB")
	}
	db.AutoMigrate(&models.Review{})
	return &ReviewRepo{DB: db}
}

func (r *ReviewRepo) AddReview(review *models.Review) error {
	return r.DB.Create(review).Error
}

func (r *ReviewRepo) GetReviews(productID string) ([]models.Review, error) {
	var reviews []models.Review
	err := r.DB.Where("product_id = ?", productID).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepo) UpdateReview(id string, rating int, comment string) error {
	return r.DB.Model(&models.Review{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"rating": rating, "comment": comment}).Error
}

func (r *ReviewRepo) DeleteReview(id string) error {
	return r.DB.Delete(&models.Review{}, "id = ?", id).Error
}
