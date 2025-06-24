package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/zoyakhan1004/e-commerce-microservices/Backend/services/user-service/internal/models"
	"github.com/zoyakhan1004/e-commerce-microservices/services/user-service/internal/repository"
	"github.com/zoyakhan1004/e-commerce-microservices/services/user-service/internal/service"
	"github.com/zoyakhan1004/e-commerce-microservices/services/user-service/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Get database configuration from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "users")
	dbPort := getEnv("DB_PORT", "5432")

	// Database connection with retry logic
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	var db *gorm.DB
	var err error

	// Retry connection up to 10 times with exponential backoff
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database (attempt %d/10): %v", i+1, err)
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after 10 attempts: %v", err)
	}

	log.Println("Successfully connected to database")

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")

	// Create repository and service
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	// gRPC server
	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, userService)

	// Start server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("User Service running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
