package models

import "time"

type User struct {
    ID        string    `gorm:"primaryKey"`
    Name      string
    Email     string    `gorm:"uniqueIndex"`
    Password  string
    Role      string
    CreatedAt time.Time
}

type PasswordReset struct {
    ID         string    `gorm:"primaryKey"`
    UserID     string
    Token      string    `gorm:"uniqueIndex"`
    ExpiresAt  time.Time
    Used       bool
}
