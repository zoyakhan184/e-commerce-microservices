package repository

import (
	"errors"
	"fmt"
	"log"
	"os"
	"user-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresUserRepo struct {
	DB *gorm.DB
}

func NewPostgresUserRepo() *PostgresUserRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// ✅ AutoMigrate your models
	err = db.AutoMigrate(
		&models.Profile{},
		&models.Address{},
		&models.Wishlist{},
	)
	if err != nil {
		log.Fatalf("failed to migrate DB: %v", err)
	}

	return &PostgresUserRepo{DB: db}
}
func (r *PostgresUserRepo) CreateProfile(p *models.Profile) error {
	return r.DB.Create(p).Error
}

func (r *PostgresUserRepo) GetProfile(userID string) (*models.Profile, error) {
	var p models.Profile
	err := r.DB.Where("user_id = ?", userID).First(&p).Error
	return &p, err
}

func (r *PostgresUserRepo) UpdateProfile(userID string, updated *models.Profile) error {
	if userID == "" {
		return fmt.Errorf("userID cannot be empty")
	}

	// Use map to avoid updating fields to zero-value unintentionally
	updateFields := map[string]interface{}{}

	if updated.FullName != "" {
		updateFields["full_name"] = updated.FullName
	}
	if updated.Phone != "" {
		updateFields["phone"] = updated.Phone
	}
	if updated.Gender != "" {
		updateFields["gender"] = updated.Gender
	}
	if updated.DOB != "" {
		updateFields["dob"] = updated.DOB
	}
	if updated.AvatarURL != "" {
		updateFields["avatar_url"] = updated.AvatarURL
	}

	// Only run update if there is something to update
	if len(updateFields) == 0 {
		log.Printf("[UpdateProfile] No fields to update for user_id=%s", userID)
		return nil
	}

	return r.DB.Model(&models.Profile{}).
		Where("user_id = ?", userID).
		Updates(updateFields).Error
}

func (r *PostgresUserRepo) AddAddress(a *models.Address) error {
	return r.DB.Create(a).Error
}

func (r *PostgresUserRepo) UpdateAddress(a *models.Address) error {
	return r.DB.Model(&models.Address{}).Where("id = ? AND user_id = ?", a.ID, a.UserID).Updates(a).Error
}

func (r *PostgresUserRepo) GetAddresses(userID string) ([]models.Address, error) {
	var addresses []models.Address
	err := r.DB.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *PostgresUserRepo) AddToWishlist(userID, productID string) error {
	w := models.Wishlist{UserID: userID, ProductID: productID}
	// INSERT ... ON CONFLICT DO NOTHING
	return r.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&w).Error
}

func (r *PostgresUserRepo) RemoveFromWishlist(userID, productID string) error {
	return r.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Wishlist{}).Error
}

func (r *PostgresUserRepo) GetWishlist(userID string) ([]string, error) {
	var wishlist []models.Wishlist
	var productIDs []string
	err := r.DB.Where("user_id = ?", userID).Find(&wishlist).Error
	for _, w := range wishlist {
		productIDs = append(productIDs, w.ProductID)
	}
	return productIDs, err
}

func (r *PostgresUserRepo) GetAllUsersWithProfiles() ([]models.UserWithProfile, error) {
	var results []models.UserWithProfile

	err := r.DB.
		Table("users").
		Select(`
		users.id AS user_id,
		users.name,
		users.email,
		COALESCE(users.role, 'user') AS role,
		users.created_at, -- ✅ make sure this is included
		profiles.full_name,
		profiles.phone,
		profiles.gender,
		profiles.dob,
		profiles.avatar_url
	`).
		Joins("LEFT JOIN profiles ON users.id = profiles.user_id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *PostgresUserRepo) UpsertProfile(userID string, profile *models.Profile) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var existing models.Profile
		result := tx.Where("user_id = ?", userID).First(&existing)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// No profile exists — create it
				if err := tx.Create(profile).Error; err != nil {
					return fmt.Errorf("failed to create profile: %w", err)
				}
			} else {
				return fmt.Errorf("failed to fetch existing profile: %w", result.Error)
			}
		} else {
			// Profile exists — update
			if err := tx.Model(&existing).Updates(profile).Error; err != nil {
				return fmt.Errorf("failed to update profile: %w", err)
			}
		}

		// ✅ Sync full_name to users table if provided
		if profile.FullName != "" {
			if err := tx.Model(&models.User{}).
				Where("id = ?", userID).
				Update("name", profile.FullName).Error; err != nil {
				return fmt.Errorf("failed to update name in users table: %w", err)
			}
		}

		return nil
	})
}

func (r *PostgresUserRepo) ClearDefaultAddress(userID string) error {
	return r.DB.Model(&models.Address{}).
		Where("user_id = ? AND is_default = ?", userID, true).
		Update("is_default", false).Error
}
