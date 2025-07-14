package repository

import "inventory-service/models"

type InventoryRepository interface {
	GetStock(productID, size, color string) (*models.Inventory, error)
	UpdateStock(skuID string, quantityChange int) error
	Restock(skuID string, quantityChange int) error
	GetBySKU(skuID string) (*models.Inventory, error) // 👈 ADD THIS LINE
	ListLowStock(threshold int) ([]models.Inventory, error)
}
