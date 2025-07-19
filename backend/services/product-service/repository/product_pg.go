package repository

import (
	"fmt"
	"os"
	"product-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresProductRepo struct {
	DB *gorm.DB
}

func NewPostgresProductRepo() *PostgresProductRepo {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to Product DB")
	}

	db.AutoMigrate(&models.Product{}, &models.ProductImage{}, &models.Category{})

	return &PostgresProductRepo{DB: db}
}

func (r *PostgresProductRepo) AddProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *PostgresProductRepo) UpdateProduct(product *models.Product) error {
	return r.DB.Save(product).Error
}

func (r *PostgresProductRepo) DeleteProduct(productID string) error {
	fmt.Printf("⛏️ Attempting to delete product ID: %s\n", productID)

	result := r.DB.Delete(&models.Product{}, "id = ?", productID)
	if result.Error != nil {
		fmt.Printf("❌ DB delete error: %v\n", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Println("⚠️ No rows affected — product may not exist")
		return fmt.Errorf("product not found")
	}

	fmt.Printf("✅ Deleted %d row(s)\n", result.RowsAffected)
	return nil
}

func (r *PostgresProductRepo) GetProduct(productID string) (*models.Product, error) {
	var product models.Product
	err := r.DB.Preload("Images").First(&product, "id = ?", productID).Error
	return &product, err
}

func (r *PostgresProductRepo) ListProducts(filter map[string]interface{}) ([]models.Product, error) {
	var products []models.Product
	tx := r.DB.Model(&models.Product{}).Preload("Images")
	for k, v := range filter {
		fmt.Println("filtering by")
		tx = tx.Where(k+" = ?", v)
	}
	err := tx.Find(&products).Error
	return products, err
}

// func (r *PostgresProductRepo) ListProducts(filter map[string]interface{}) ([]models.Product, error) {
// 	var products []models.Product

// 	tx := r.DB.Model(&models.Product{}).Preload("Images")
// 	for k, v := range filter {
// 		tx = tx.Where(k+" = ?", v)
// 	}

// 	err := tx.Find(&products).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Inject base64 image data for each product image
// 	for i := range products {
// 		for j := range products[i].Images {
// 			img := &products[i].Images[j]
// 			if img.FileName != "" {
// 				fullPath := filepath.Join("uploads", img.FileName)
// 				data, err := os.ReadFile(fullPath)
// 				if err != nil {
// 					continue // skip this image if file can't be read
// 				}

// 				// Infer MIME type
// 				ext := filepath.Ext(img.FileName)
// 				mimeType := "image/jpeg"
// 				switch ext {
// 				case ".png":
// 					mimeType = "image/png"
// 				case ".gif":
// 					mimeType = "image/gif"
// 				case ".webp":
// 					mimeType = "image/webp"
// 				}

// 				// Convert to base64
// 				base64Str := base64.StdEncoding.EncodeToString(data)
// 				fmt.Println("base64Str", base64Str)
// 				img.Base64Data = fmt.Sprintf("data:%s;base64,%s", mimeType, base64Str)
// 			}
// 		}
// 	}

// 	return products, nil
// }

func (r *PostgresProductRepo) AddCategory(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *PostgresProductRepo) ListCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

func (r *PostgresProductRepo) ListLowStock(threshold int) ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload("Images").Where("quantity < ?", threshold).Find(&products).Error
	return products, err
}
