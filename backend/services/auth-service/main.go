// auth-service/main.go
package main

import (
	"auth-service/handlers"
	"auth-service/models"
	"auth-service/rabbitmq"
	"auth-service/repository"
	authpb "auth-service/proto"

	"google.golang.org/grpc"

	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, relying on system env vars")
	}

	// Connect to DB using repository constructor
	repo := repository.NewPostgresUserRepo()
	db := repo.DB
	db.AutoMigrate(&models.User{}, &models.PasswordReset{})

	// Initialize RabbitMQ
	rabbitmq.InitRabbitMQ(os.Getenv("RABBITMQ_URL"))

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authpb .RegisterAuthServiceServer(grpcServer, &handlers.AuthService{DB: db})
	log.Println("Auth service listening on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
