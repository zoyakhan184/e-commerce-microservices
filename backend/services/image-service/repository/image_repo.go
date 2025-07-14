package repository

import (
	"image-service/models"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ImageRepo struct {
	DB *gorm.DB
}

func NewImageRepo() *ImageRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Failed to connect to database: " + err.Error())
	}

	if err := db.AutoMigrate(&models.Image{}); err != nil {
		panic("❌ AutoMigrate failed: " + err.Error())
	}

	return &ImageRepo{DB: db}
}

func (r *ImageRepo) Save(image *models.Image) error {
	return r.DB.Create(image).Error
}

func (r *ImageRepo) Get(id string) (*models.Image, error) {
	var img models.Image
	err := r.DB.First(&img, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &img, nil
}

func (r *ImageRepo) Delete(id string) error {
	img, err := r.Get(id)
	if err != nil {
		return err
	}

	// Delete image file from disk
	if img.FileName != "" {
		fullPath := filepath.Join("uploads", img.FileName)
		if err := os.Remove(fullPath); err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	// Delete DB record
	return r.DB.Delete(&models.Image{}, "id = ?", id).Error
}
