package models

type Profile struct {
	UserID    string `gorm:"primaryKey"`
	FullName  string
	Phone     string
	Gender    string
	DOB       string
	AvatarURL string
	Email     string `gorm:"-"` // ← Ignored by GORM during insert/update
}
