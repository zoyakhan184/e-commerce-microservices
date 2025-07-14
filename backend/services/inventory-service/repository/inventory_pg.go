package repository

import (
	"inventory-service/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresInventoryRepo struct {
	DB *gorm.DB
}

const LowStockThreshold = 5

func NewPostgresInventoryRepo() *PostgresInventoryRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Failed to connect to DB: " + err.Error())
	}

	db.AutoMigrate(&models.Inventory{})
	return &PostgresInventoryRepo{DB: db}
}

func (r *PostgresInventoryRepo) GetStock(productID, size, color string) (*models.Inventory, error) {
	var inv models.Inventory
	err := r.DB.Where("product_id = ? AND size = ? AND color = ?", productID, size, color).First(&inv).Error
	return &inv, err
}

func (r *PostgresInventoryRepo) GetBySKU(skuID string) (*models.Inventory, error) {
	var inv models.Inventory
	err := r.DB.Where("sku_id = ?", skuID).First(&inv).Error
	return &inv, err
}

func (r *PostgresInventoryRepo) UpdateStock(skuID string, quantityChange int) error {
	err := r.DB.Model(&models.Inventory{}).
		Where("sku_id = ?", skuID).
		Update("quantity", gorm.Expr("quantity + ?", quantityChange)).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresInventoryRepo) Restock(skuID string, quantityChange int) error {
	return r.UpdateStock(skuID, quantityChange)
}

func (r *PostgresInventoryRepo) ListLowStock(threshold int) ([]models.Inventory, error) {
	var list []models.Inventory
	err := r.DB.Where("quantity <= ?", threshold).Find(&list).Error
	return list, err
}
