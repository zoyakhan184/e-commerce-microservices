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

	// Enable CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // adjust for your frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")

	// Auth
	api.POST("/auth/login", handlers.Login)
	api.POST("/auth/register", handlers.Register)
	api.GET("/auth/validate", handlers.ValidateToken)
	api.POST("/auth/forgot", handlers.ForgotPassword)
	api.POST("/auth/reset", handlers.ResetPassword)

	// Users
	userGroup := api.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	userGroup.GET("/profile", handlers.GetUserProfile)
	userGroup.PUT("/profile", handlers.UpdateUserProfile)
	userGroup.POST("/address", handlers.AddAddress)
	userGroup.GET("/addresses", handlers.GetAddresses)
	userGroup.POST("/wishlist", handlers.AddToWishlist)
	userGroup.GET("/wishlist", handlers.GetWishlist)
	userGroup.DELETE("/wishlist/:productId", handlers.RemoveFromWishlist)

	// Products
	api.GET("/products/:id", handlers.GetProduct)
	api.GET("/products", handlers.ListProducts)
	api.POST("/products", handlers.AddProduct)
	api.PUT("/products/:id", handlers.EditProduct)
	api.DELETE("/products/:id", handlers.DeleteProduct)

	// Categories
	api.GET("/categories", handlers.ListCategories)
	api.POST("/categories", handlers.AddCategory)

	// Inventory
	api.GET("/inventory/stock", handlers.GetStock)
	api.POST("/inventory/stock/update", handlers.UpdateStockOnOrder) // 🔧 fixed missing "/"
	api.POST("/inventory/stock/restock", handlers.Restock)           // 🔧 fixed missing "/"
	api.GET("/inventory/low-stock", handlers.ListLowStock)           // 🔧 fixed missing "/"

	// Cart
	api.Use(middleware.AuthMiddleware()) // ✅ MUST include this!

	{
		api.POST("/cart", handlers.AddToCart)
		api.GET("/cart", handlers.GetCart)
		api.PUT("/cart", handlers.UpdateCartItem)
		api.DELETE("/cart", handlers.RemoveFromCart)
	}

	// Orders
	api.POST("/orders", handlers.PlaceOrder)
	api.GET("/orders", handlers.GetOrders)
	api.GET("/orders/:id", handlers.GetOrderDetails)
	api.PUT("/orders/:id/status", handlers.UpdateOrderStatus)

	// Payments
	api.POST("/payment/initiate", handlers.InitiatePayment)
	api.POST("/payment/verify", handlers.VerifyPayment)
	api.POST("/payment/refund", handlers.RefundPayment)

	// Reviews
	api.POST("/reviews", handlers.AddReview)
	api.GET("/reviews/:product_id", handlers.GetReviews)
	api.PUT("/reviews/:id", handlers.EditReview)
	api.DELETE("/reviews/:id", handlers.DeleteReview)

	// Images
	api.POST("/images/upload", handlers.UploadImage)
	api.DELETE("/images/delete", handlers.DeleteImage)
	api.GET("/images/:id", handlers.GetImage) // 🔧 fixed route clash (was just "/:id")

	// Admin
	api.GET("/admin/dashboard", handlers.GetDashboard)
	api.GET("/admin/users", handlers.ListUsers)
	api.GET("/admin/orders", handlers.ListAllOrders)
	api.GET("/admin/activity", handlers.GetRecentActivity)

	api.GET("/admin/user-profiles", handlers.ListAllUserProfiles)

	return r
}
