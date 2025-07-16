package router

import (
	"bff-service/handlers"
	"bff-service/middleware" // ✅ Import your middleware

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")

	// Auth (public)
	api.POST("/auth/login", handlers.Login)
	api.POST("/auth/register", handlers.Register)
	api.GET("/auth/validate", handlers.ValidateToken)
	api.POST("/auth/forgot", handlers.ForgotPassword)
	api.POST("/auth/reset", handlers.ResetPassword)

	// Users (protected)
	userGroup := api.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/profile", handlers.GetUserProfile)
		userGroup.PUT("/profile", handlers.UpdateUserProfile)
		userGroup.POST("/address", handlers.AddAddress)
		userGroup.GET("/addresses", handlers.GetAddresses)
		userGroup.POST("/wishlist", handlers.AddToWishlist)
		userGroup.GET("/wishlist", handlers.GetWishlist)
		userGroup.DELETE("/wishlist/:productId", handlers.RemoveFromWishlist)
	}

	// Cart (protected)
	cartGroup := api.Group("/cart")
	cartGroup.Use(middleware.AuthMiddleware())
	{
		cartGroup.POST("", handlers.AddToCart)
		cartGroup.GET("", handlers.GetCart)
		cartGroup.PUT("", handlers.UpdateCartItem)
		cartGroup.DELETE("", handlers.RemoveFromCart)
		cartGroup.DELETE("/clear", handlers.ClearCart)
	}

	// Orders (protected)
	orderGroup := api.Group("/orders")
	orderGroup.Use(middleware.AuthMiddleware())
	{
		orderGroup.POST("", handlers.PlaceOrder)
		orderGroup.GET("", handlers.GetOrders)
		orderGroup.GET("/:id", handlers.GetOrderDetails)
		orderGroup.PUT("/:id/status", handlers.UpdateOrderStatus)
		//orderGroup.POST("", handlers.ListAllOrders)
	}

	// Reviews (protected)
	reviewGroup := api.Group("/reviews")
	reviewGroup.Use(middleware.AuthMiddleware())
	{
		reviewGroup.POST("", handlers.AddReview)
		reviewGroup.GET("/:product_id", handlers.GetReviews)
		reviewGroup.PUT("/:id", handlers.EditReview)
		reviewGroup.DELETE("/:id", handlers.DeleteReview)
	}

	// Images (public)
	api.POST("/images/upload", handlers.UploadImage)
	api.DELETE("/images/delete", handlers.DeleteImage)
	api.GET("/images/:id", handlers.GetImage)

	// Products (public)
	api.GET("/products/:id", handlers.GetProduct)
	api.GET("/products", handlers.ListProducts)
	api.POST("/products", handlers.AddProduct)
	api.PUT("/products/:id", handlers.EditProduct)
	api.DELETE("/products/:id", handlers.DeleteProduct)

	// Categories (public)
	api.GET("/categories", handlers.ListCategories)
	api.POST("/categories", handlers.AddCategory)

	// Inventory (protected)
	invGroup := api.Group("/inventory")
	invGroup.Use(middleware.AuthMiddleware())
	{
		invGroup.GET("/stock", handlers.GetStock)
		invGroup.POST("/stock/update", handlers.UpdateStockOnOrder)
		invGroup.POST("/stock/restock", handlers.Restock)
		invGroup.GET("/low-stock", handlers.ListLowStock)
	}

	// Admin (should be protected with admin-only middleware)
	adminGroup := api.Group("/admin")
	adminGroup.Use(middleware.AuthMiddleware()) // ❗ optionally: .Use(middleware.AdminOnly())
	{
		adminGroup.GET("/dashboard", handlers.GetDashboard)
		adminGroup.GET("/users", handlers.ListUsers)
		adminGroup.GET("/orders", handlers.ListAllOrders)
		adminGroup.GET("/activity", handlers.GetRecentActivity)
		adminGroup.GET("/user-profiles", handlers.ListAllUserProfiles)
	}

	return r
}
