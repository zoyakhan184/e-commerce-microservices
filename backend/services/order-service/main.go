package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"

	"order-service/handlers"
	orderpb "order-service/proto"
	"order-service/rabbitmq"
	"order-service/repository"

	productpb "product-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	_ = godotenv.Load("../.env")

	// Setup database
	repo := repository.NewOrderRepo()

	// Setup RabbitMQ
	rabbitmq.InitRabbitMQ()
	rabbitmq.StartConsumer(repo)

	// Get product service address with fallback
	productServiceAddr := os.Getenv("PRODUCT_SERVICE_ADDR")
	if productServiceAddr == "" {
		productServiceAddr = "localhost:50059" // Default to match your product service port
		log.Printf("⚠️  PRODUCT_SERVICE_ADDR not set, using default: %s", productServiceAddr)
	}

	log.Printf("🔗 Connecting to product service at: %s", productServiceAddr)

	// Connect to product-service (updated to use new gRPC credentials API)
	productConn, err := grpc.Dial(
		productServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("❌ Failed to connect to product-service: %v", err)
	}
	defer productConn.Close()

	log.Println("✅ Connected to product service")

	productClient := productpb.NewProductServiceClient(productConn)

	// gRPC server setup
	lis, err := net.Listen("tcp", ":50057")
	if err != nil {
		log.Fatalf("❌ Failed to listen on port 50057: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, &handlers.OrderService{
		Repo:          repo,
		ProductClient: productClient,
	})

	log.Println("📦 Order Service running on port 50057")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
