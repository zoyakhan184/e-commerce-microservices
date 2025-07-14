package models

type Address struct {
	ID          string `gorm:"primaryKey;column:id"`
	UserID      string `gorm:"column:user_id"`
	Name        string `gorm:"column:name"`
	Phone       string `gorm:"column:phone"`
	AddressLine string `gorm:"column:address_line"`
	City        string `gorm:"column:city"`
	State       string `gorm:"column:state"`
	Zip         string `gorm:"column:zip"`
	Country     string `gorm:"column:country"`
	IsDefault   bool   `gorm:"column:is_default"`
}
