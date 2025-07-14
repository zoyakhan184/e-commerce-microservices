package repository

import "user-service/models"

type UserRepository interface {
	CreateProfile(profile *models.Profile) error
	GetProfile(userID string) (*models.Profile, error)
	UpdateProfile(userID string, updated *models.Profile) error
	UpsertProfile(userID string, profile *models.Profile) error
	GetAllUsersWithProfiles() ([]models.UserWithProfile, error)

	AddAddress(addr *models.Address) error
	UpdateAddress(addr *models.Address) error
	GetAddresses(userID string) ([]models.Address, error)

	// ✅ NEW
	ClearDefaultAddress(userID string) error

	AddToWishlist(userID, productID string) error
	RemoveFromWishlist(userID, productID string) error
	GetWishlist(userID string) ([]string, error)
}
