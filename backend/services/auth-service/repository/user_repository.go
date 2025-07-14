package repository

import (
    "auth-service/models"
)

type UserRepository interface {
    CreateUser(user *models.User) error
    GetUserByEmail(email string) (*models.User, error)
    UpdatePassword(userID, newHash string) error
}
