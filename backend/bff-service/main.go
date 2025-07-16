package main

import (
	"bff-service/clients"
	"bff-service/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// ✅ Load .env from current directory (same as main.go)
	if err := godotenv.Load(".env"); err != nil {

		log.Println("⚠️  .env file not found in root, using system environment variables")
	} else {
		log.Println("✅ .env file loaded successfully from root")
	}

	// ✅ Log JWT_SECRET presence
	jwtSecret := os.Getenv("JWT_SECRET")
	fmt.Println("JWT_SECRET:", jwtSecret)
	if jwtSecret == "" {
		log.Println("❌ JWT_SECRET not set. Token validation will fail.")
	} else {
		log.Printf("🔐 JWT_SECRET is set (length: %d)\n", len(jwtSecret))
	}

	// Initialize gRPC clients
	log.Println("🔌 Initializing gRPC clients...")
	clients.InitImageClient()
	log.Println("✅ ImageClient initialized")
	clients.InitReviewClient()
	log.Println("✅ ReviewClient initialized")
	clients.InitInventoryClient()
	log.Println("✅ InventoryClient initialized")
	clients.AdminClient()
	log.Println("✅ AdminClient initialized")
	clients.AuthClient()
	log.Println("✅ AuthClient initialized")
	clients.OrderClient()
	log.Println("✅ OrderClient initialized")
	clients.PaymentClient()
	log.Println("✅ PaymentClient initialized")
	clients.ProductClient()
	log.Println("✅ ProductClient initialized")
	clients.CartClient()
	log.Println("✅ CartClient initialized")

	// Start Gin server
	r := router.SetupRouter()
	port := os.Getenv("BFF_PORT")
	if port == "" {
		port = "8080"
		log.Println("ℹ️  BFF_PORT not set, defaulting to :8080")
	}
	log.Printf("🚀 BFF Server running at http://localhost:%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
