package models

type Image struct {
	ID         string `gorm:"primaryKey"`
	EntityID   string
	EntityType string
	FileType   string
	FileName   string // ✅ Store only file name like "abc.jpg"
}
