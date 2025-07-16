package handlers

import (
	"bff-service/clients"
	productpb "bff-service/proto/product"
	userpb "bff-service/proto/user"
	"bff-service/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ---------- Profile ----------

func GetUserProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	log.Printf("[GetUserProfile] user_id=%s", userID)

	resp, err := clients.UserClient().GetUser(c, &userpb.GetUserRequest{UserId: userID})
	if err != nil {
		log.Printf("[GetUserProfile] Failed to fetch profile for user_id=%s: %v", userID, err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch user profile")
		return
	}
	log.Printf("[GetUserProfile] Success for user_id=%s", userID)
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

func UpdateUserProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		log.Printf("[UpdateUserProfile] ❌ Missing user_id in context")
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: user_id not found")
		return
	}
	log.Printf("[UpdateUserProfile] 🔐 Authenticated user_id=%s", userID)

	// Parse JSON input
	var input struct {
		FullName  string `json:"full_name"`
		Phone     string `json:"phone"`
		Gender    string `json:"gender"`
		Dob       string `json:"dob"`
		AvatarUrl string `json:"avatar_url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("[UpdateUserProfile] ⚠️ Invalid input: %v", err)
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	log.Printf("[UpdateUserProfile] 📥 Received update payload: %+v", input)

	userProfile := &userpb.UserProfile{
		UserId:    userID,
		FullName:  input.FullName,
		Phone:     input.Phone,
		Gender:    input.Gender,
		Dob:       input.Dob,
		AvatarUrl: input.AvatarUrl,
	}

	// Attempt gRPC call to update or create profile
	resp, err := clients.UserClient().UpdateUser(c, userProfile)
	if err != nil {
		log.Printf("[UpdateUserProfile] ❌ Failed to update profile for user_id=%s: %v", userID, err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update user profile")
		return
	}

	log.Printf("[UpdateUserProfile] ✅ Success - user_id=%s", userID)
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

func AddAddress(c *gin.Context) {
	userID := c.GetString("user_id")
	log.Printf("[AddAddress] 👤 user_id=%s", userID)

	var req userpb.AddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[AddAddress] ⚠️ Invalid input: %v", err)
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid address input")
		return
	}
	req.UserId = userID

	log.Printf("[AddAddress] 📦 Payload: %+v", req)

	resp, err := clients.UserClient().AddAddress(c, &req)
	if err != nil {
		log.Printf("[AddAddress] ❌ Failed: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to add address")
		return
	}

	log.Printf("[AddAddress] ✅ Success: %+v", resp)
	utils.RespondWithJSON(c, http.StatusOK, gin.H{"address": resp})
}

func UpdateAddress(c *gin.Context) {
	userID := c.GetString("user_id")
	log.Printf("[UpdateAddress] 👤 user_id=%s", userID)

	var req userpb.AddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[UpdateAddress] ⚠️ Invalid input: %v", err)
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid address input")
		return
	}
	req.UserId = userID

	log.Printf("[UpdateAddress] 📦 Payload: %+v", req)

	resp, err := clients.UserClient().UpdateAddress(c, &req)
	if err != nil {
		log.Printf("[UpdateAddress] ❌ Failed: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update address")
		return
	}

	log.Printf("[UpdateAddress] ✅ Success: %+v", resp)
	utils.RespondWithJSON(c, http.StatusOK, gin.H{"address": resp})
}

func GetAddresses(c *gin.Context) {
	userID := c.GetString("user_id")
	log.Printf("[GetAddresses] 🔍 user_id=%s", userID)

	resp, err := clients.UserClient().GetAddresses(c, &userpb.UserRequest{UserId: userID})
	if err != nil {
		log.Printf("[GetAddresses] ❌ Failed: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch addresses")
		return
	}

	log.Printf("[GetAddresses] ✅ Fetched %d addresses", len(resp.Addresses))
	utils.RespondWithJSON(c, http.StatusOK, gin.H{"addresses": resp.Addresses})
}

// ---------- Wishlist ----------

func AddToWishlist(c *gin.Context) {
	userID := c.GetString("user_id")
	log.Printf("[AddToWishlist] user_id=%s", userID)

	var req userpb.WishlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[AddToWishlist] Invalid input: %v", err)
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid wishlist input")
		return
	}

	if req.ProductId == "" {
		log.Printf("[AddToWishlist] Missing product_id for user_id=%s", userID)
		utils.RespondWithError(c, http.StatusBadRequest, "Product ID is required")
		return
	}

	req.UserId = userID

	resp, err := clients.UserClient().AddToWishlist(c, &req)
	if err != nil {
		log.Printf("[AddToWishlist] Failed to add to wishlist for user_id=%s: %v", userID, err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to add to wishlist")
		return
	}

	log.Printf("[AddToWishlist] Success for user_id=%s, product_id=%s", userID, req.ProductId)
	utils.RespondWithJSON(c, http.StatusOK, resp)
}

func RemoveFromWishlist(c *gin.Context) {
	userID := c.GetString("user_id")
	productID := c.Param("productId") // ✅ get from URL path

	log.Printf("[RemoveFromWishlist] user_id=%s, product_id=%s", userID, productID)

	if productID == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Product ID is required")
		return
	}

	resp, err := clients.UserClient().RemoveFromWishlist(c, &userpb.WishlistRequest{
		UserId:    userID,
		ProductId: productID,
	})
	if err != nil {
		log.Printf("[RemoveFromWishlist] Failed to remove from wishlist for user_id=%s: %v", userID, err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to remove from wishlist")
		return
	}

	log.Printf("[RemoveFromWishlist] Success for user_id=%s, product_id=%s", userID, productID)
	utils.RespondWithJSON(c, http.StatusOK, resp)
}
func GetWishlist(c *gin.Context) {
	userID := c.GetString("user_id")
	log.Printf("[GetWishlist] user_id=%s", userID)

	// 1. Fetch wishlist product IDs from user-service
	resp, err := clients.UserClient().GetWishlist(c, &userpb.UserRequest{UserId: userID})
	if err != nil {
		log.Printf("[GetWishlist] ❌ Failed to fetch wishlist: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch wishlist")
		return
	}

	// WishlistItem represents a product in the user's wishlist
	type WishlistItem struct {
		ProductID   string  `json:"product_id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
		ImageUrl    string  `json:"image_url"`
	}
	
	// 2. Build enriched product data
		var enriched []*WishlistItem
		for _, productID := range resp.ProductIds {
			productResp, err := clients.ProductClient().GetProduct(c, &productpb.ProductIdRequest{Id: productID})
			if err != nil {
				log.Printf("[GetWishlist] ⚠️ Failed to fetch product %s: %v", productID, err)
				continue
			}
	
			var imageUrl string
			if len(productResp.ImageUrls) > 0 {
				imageUrl = fmt.Sprintf("data:image/jpeg;base64,%s", productResp.ImageUrls[0])
			}
	
			enriched = append(enriched, &WishlistItem{
				ProductID:   productResp.Id,
				Name:        productResp.Name,
				Description: productResp.Description,
				Price:       float32(productResp.Price),
				ImageUrl:    imageUrl,
			})
		}

	// 3. Respond with full wishlist
	utils.RespondWithJSON(c, http.StatusOK, gin.H{
		"items": enriched,
	})
}

func ListAllUserProfiles(c *gin.Context) {
	log.Println("[ListAllUserProfiles] Called")

	resp, err := clients.UserClient().ListAllProfiles(c, &userpb.Empty{})
	if err != nil {
		log.Printf("[ListAllUserProfiles] ❌ Failed to list profiles: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to list user profiles")
		return
	}

	// Optionally, you can transform or validate data here
	log.Printf("[ListAllUserProfiles] ✅ Success: %d profiles retrieved", len(resp.Profiles))
	utils.RespondWithJSON(c, http.StatusOK, resp.Profiles) // unwrap `.Profiles` to directly return list
}
