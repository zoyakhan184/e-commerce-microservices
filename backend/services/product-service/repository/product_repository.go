package repository

import "product-service/models"

type ProductRepository interface {
	AddProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(productID string) error
	GetProduct(productID string) (*models.Product, error)
	ListProducts(filter map[string]interface{}) ([]models.Product, error)

	AddCategory(category *models.Category) error
	ListCategories() ([]models.Category, error)
}
