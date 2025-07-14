package models

type Category struct {
    ID       string `gorm:"primaryKey"`
    Name     string
    Gender   string // Men, Women, Kids
    ParentID *string // nullable for subcategories
}
