package handlers

import (
	"cart-service/models"
	cartpb "cart-service/proto"
	"cart-service/repository"
	"context"
	"fmt"
	"log"
	"time"
)

type CartService struct {
	Repo repository.CartRepository
	cartpb.UnimplementedCartServiceServer
}

// AddToCart adds an item to the user's cart.
func (s *CartService) AddToCart(ctx context.Context, req *cartpb.CartItem) (*cartpb.GenericResponse, error) {
	log.Printf("[AddToCart] Called for user %s: ProductID=%s, Size=%s, Color=%s, Qty=%d",
		req.UserId, req.ProductId, req.Size, req.Color, req.Quantity)

	item := &models.CartItem{
		UserID:    req.UserId,
		ProductID: req.ProductId,
		Size:      req.Size,
		Color:     req.Color,
		Quantity:  int(req.Quantity),
		AddedAt:   time.Now(),
	}

	if err := s.Repo.AddToCart(item); err != nil {
		log.Printf("[AddToCart] ❌ Failed to add to cart: %v", err)
		return &cartpb.GenericResponse{Message: "Failed to add to cart"}, err
	}

	log.Printf("[AddToCart] ✅ Item added to cart for user %s", req.UserId)
	return &cartpb.GenericResponse{Message: "Item added to cart"}, nil
}

// RemoveFromCart deletes a product from the user's cart.
func (s *CartService) RemoveFromCart(ctx context.Context, req *cartpb.CartRequest) (*cartpb.GenericResponse, error) {
	log.Printf("[RemoveFromCart] Called for user %s: ProductID=%s", req.UserId, req.ProductId)

	if err := s.Repo.RemoveFromCart(req.UserId, req.ProductId); err != nil {
		log.Printf("[RemoveFromCart] ❌ Failed to remove item: %v", err)
		return &cartpb.GenericResponse{Message: "Failed to remove item"}, err
	}

	log.Printf("[RemoveFromCart] ✅ Item removed from cart for user %s", req.UserId)
	return &cartpb.GenericResponse{Message: "Item removed from cart"}, nil
}

// UpdateCartItem updates the quantity of a cart item.
func (s *CartService) UpdateCartItem(ctx context.Context, req *cartpb.CartUpdateRequest) (*cartpb.GenericResponse, error) {
	log.Printf("[UpdateCartItem] Called for user %s: ProductID=%s, NewQty=%d", req.UserId, req.ProductId, req.Quantity)

	if err := s.Repo.UpdateCartItem(req.UserId, req.ProductId, int(req.Quantity)); err != nil {
		log.Printf("[UpdateCartItem] ❌ Failed to update item: %v", err)
		return &cartpb.GenericResponse{Message: "Failed to update item"}, err
	}

	log.Printf("[UpdateCartItem] ✅ Item updated for user %s", req.UserId)
	return &cartpb.GenericResponse{Message: "Item updated"}, nil
}

// GetCart fetches the list of items in a user's cart.
func (s *CartService) GetCart(ctx context.Context, req *cartpb.UserRequest) (*cartpb.CartList, error) {
	log.Printf("[GetCart] Called for user %s", req.UserId)

	items, err := s.Repo.GetCart(req.UserId)
	if err != nil {
		log.Printf("[GetCart] ❌ Failed to fetch cart: %v", err)
		return nil, err
	}

	log.Printf("[GetCart] ✅ Retrieved %d cart items for user %s", len(items), req.UserId)

	var cart []*cartpb.CartItem
	for _, item := range items {
		fmt.Println("Cart Items tmkc:", item)
		cart = append(cart, &cartpb.CartItem{
			UserId:    item.UserID,
			ProductId: item.ProductID,
			Size:      item.Size,
			Color:     item.Color,
			Quantity:  int32(item.Quantity),
		})
		fmt.Print("cart:", cart)
	}

	return &cartpb.CartList{Items: cart}, nil
}

func (s *CartService) ClearCart(ctx context.Context, req *cartpb.UserRequest) (*cartpb.GenericResponse, error) {
	log.Printf("[ClearCart] Called for user %s", req.UserId)

	if err := s.Repo.ClearCart(req.UserId); err != nil {
		log.Printf("[ClearCart] ❌ Failed to clear cart: %v", err)
		return &cartpb.GenericResponse{Message: "Failed to clear cart"}, err
	}

	log.Printf("[ClearCart] ✅ Cleared cart for user %s", req.UserId)
	return &cartpb.GenericResponse{Message: "Cart cleared"}, nil
}
