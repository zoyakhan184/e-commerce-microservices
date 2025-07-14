package models

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string
	Role  string `gorm:"default:user"` // only applies on insert
}
