package main

import (
	"admin-service/handlers"
	adminpb "admin-service/proto"
	"admin-service/rabbitmq"
	"admin-service/repository"
	"log"
	"net"
	"os"

	userpb "user-service/proto" // 👈 Import the user-service proto package

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load("../.env")

	repo := repository.NewAdminRepo()

	// ✅ Start all RabbitMQ listeners
	rabbitmq.StartListeners(repo)

	// 🔗 Connect to user-service via gRPC
	userServiceAddress := os.Getenv("USER_SERVICE_ADDR") // e.g. "user-service:50052"
	if userServiceAddress == "" {
		userServiceAddress = "localhost:50052"
	}

	conn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}
	defer conn.Close()

	userClient := userpb.NewUserServiceClient(conn)

	// 🛠️ Create AdminService handler with repo + gRPC client
	adminService := &handlers.AdminService{
		Repo:       repo,
		UserClient: userClient,
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	adminpb.RegisterAdminServiceServer(grpcServer, adminService)

	log.Println("🧑‍💼 Admin Service running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
