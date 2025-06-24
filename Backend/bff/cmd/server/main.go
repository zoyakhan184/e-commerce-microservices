package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zoyakhan1004/e-commerce-microservices/bff/internal/clients"
	"github.com/zoyakhan1004/e-commerce-microservices/bff/internal/handlers"
)

func main() {
	// Initialize gRPC clients
	grpcClients, err := clients.NewGrpcClients()
	if err != nil {
		log.Fatalf("failed to initialize gRPC clients: %v", err)
	}

	// Create HTTP server
	r := gin.Default()

	// Register routes
	userHandler := handlers.NewUserHandler(grpcClients)
	r.POST("/api/auth/register", userHandler.Register)
	r.POST("/api/auth/login", userHandler.Login)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Start server
	log.Println("BFF server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
