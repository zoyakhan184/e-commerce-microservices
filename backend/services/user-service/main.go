// user-service/main.go
package main

import (
	"log"
	"net"

	"user-service/handlers"
	userpb "user-service/proto"
	"user-service/repository"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, relying on system env vars")
	}

	// Initialize repository with DB connection from env
	repo := repository.NewPostgresUserRepo()
	repository.InitGlobalRepo(repo)
	userHandler := &handlers.UserService{Repo: repo}

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50061")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Println("User Service is running on port :50061")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
