package handlers

import (
	"bff-service/clients"
	cartpb "bff-service/proto/cart"
	productpb "bff-service/proto/product"
	"bff-service/utils"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartSummaryResponse struct {
	Items    []*cartpb.CartItem `json:"items"`
	Subtotal float32            `json:"subtotal"`
	Shipping float32            `json:"shipping"`
	Total    float32            `json:"total"`
}

func AddToCart(c *gin.Context) {
	log.Println("[AddToCart] Called")

	uid := c.GetString("user_id")
	var req cartpb.CartItem
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[AddToCart] ❌ Failed to bind JSON: %v", err)
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	req.UserId = uid
	log.Printf("[AddToCart] ➕ Adding product %s (qty: %d) for user %s", req.ProductId, req.Quantity, uid)

	resp, err := clients.CartClient().AddToCart(c, &req)
	if err != nil {
		log.Printf("[AddToCart] ❌ gRPC error: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Error")
		return
	}

	log.Printf("[AddToCart] ✅ Successfully added item for user %s", uid)
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

func GetCart(c *gin.Context) {
	log.Println("[GetCart] Called")

	uid := c.GetString("user_id")
	log.Printf("[GetCart] 🛒 Fetching cart for user %s", uid)

	cartResp, err := clients.CartClient().GetCart(c, &cartpb.UserRequest{UserId: uid})
	if err != nil {
		log.Printf("[GetCart] ❌ gRPC error: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Error fetching cart")
		return
	}
	fmt.Println("hehhe", cartResp.Items)
	var subtotal float32
	for _, item := range cartResp.Items {
		productResp, err := clients.ProductClient().GetProduct(c, &productpb.ProductIdRequest{Id: item.ProductId})
		if err != nil {
			log.Printf("[GetCart] ⚠️ Failed to fetch product %s: %v", item.ProductId, err)
			continue
		}

		item.ProductName = productResp.Name
		item.Price = float32(productResp.Price)
		fmt.Println("Product Name:", item.ProductName, "Price:", item.Price)
		if len(productResp.ImageUrls) > 0 {
			fmt.Println("Image URLs:", productResp.ImageUrls)
			dataURI := fmt.Sprintf("data:image/jpeg;base64,%s", productResp.ImageUrls[0])
			item.ImageUrl = dataURI
		}

		subtotal += item.Price * float32(item.Quantity)
	}

	shipping := float32(10.0)
	total := subtotal + shipping

	response := CartSummaryResponse{
		Items:    cartResp.Items,
		Subtotal: subtotal,
		Shipping: shipping,
		Total:    total,
	}

	utils.RespondWithJSON(c, http.StatusOK, response)
}

func RemoveFromCart(c *gin.Context) {
	log.Println("[RemoveFromCart] Called")

	uid := c.GetString("user_id")
	var req cartpb.CartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[RemoveFromCart] ❌ Failed to bind JSON: %v", err)
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	req.UserId = uid

	log.Printf("[RemoveFromCart] 🗑️ Removing product %s for user %s", req.ProductId, uid)

	resp, err := clients.CartClient().RemoveFromCart(c, &req)
	if err != nil {
		log.Printf("[RemoveFromCart] ❌ gRPC error: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to remove item")
		return
	}

	log.Printf("[RemoveFromCart] ✅ Removed item for user %s", uid)
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

func UpdateCartItem(c *gin.Context) {
	log.Println("[UpdateCartItem] Called")

	uid := c.GetString("user_id")
	var req cartpb.CartUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[UpdateCartItem] ❌ Failed to bind JSON: %v", err)
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	req.UserId = uid

	log.Printf("[UpdateCartItem] 🔁 Updating product %s (qty: %d) for user %s", req.ProductId, req.Quantity, uid)

	resp, err := clients.CartClient().UpdateCartItem(c, &req)
	if err != nil {
		log.Printf("[UpdateCartItem] ❌ gRPC error: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update item")
		return
	}

	log.Printf("[UpdateCartItem] ✅ Updated cart item for user %s", uid)
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

func ClearCart(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	req := &cartpb.UserRequest{UserId: userID}
	_, err := clients.CartClient().ClearCart(context.Background(), req)
	if err != nil {
		log.Printf("[ClearCart] ❌ gRPC error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}

	log.Printf("[ClearCart] ✅ Cleared cart for user %s", userID)
	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared successfully"})
}
