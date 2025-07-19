package models

import "time"

type UserWithProfile struct {
	UserID    string
	Name      string
	Role      string
	Email     string
	FullName  string
	Phone     string
	Gender    string
	DOB       string
	AvatarURL string
	CreatedAt time.Time // ✅ Add this field
}
