package repository

import "cart-service/models"

type CartRepository interface {
	AddToCart(item *models.CartItem) error
	RemoveFromCart(userID, productID string) error
	UpdateCartItem(userID, productID string, quantity int) error
	GetCart(userID string) ([]models.CartItem, error)
	ClearCart(userID string) error // ✅ New method

}
