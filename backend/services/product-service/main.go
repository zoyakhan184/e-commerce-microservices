package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"

	"product-service/handlers"
	productpb "product-service/proto"
	"product-service/rabbitmq"
	"product-service/repository"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("⚠️ No .env file found, using system environment")
	}

	// ✅ Connect to RabbitMQ
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ Failed to open RabbitMQ channel: %v", err)
	}
	defer channel.Close()

	// ✅ Declare low-stock queue
	_, err = channel.QueueDeclare("product.low_stock", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("❌ Failed to declare queue: %v", err)
	}

	// ✅ Pass RabbitMQ channel to publisher
	rabbitmq.Init(channel)

	// Init product repository & handler
	repo := repository.NewPostgresProductRepo()
	handler := &handlers.ProductService{Repo: repo}

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50059")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	productpb.RegisterProductServiceServer(grpcServer, handler)

	log.Println("🚀 Product Service running on :50059")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
