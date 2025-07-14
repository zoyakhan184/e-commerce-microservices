package main

import (
	"log"
	"net"
	"os"
	"payment-service/gateway"
	"payment-service/handlers"
	paymentpb "payment-service/proto"
	"payment-service/rabbitmq"
	"payment-service/repository"
	"payment-service/webhook"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("⚠️  No .env file found. Relying on system env vars.")
	}

	// Check essential env variables
	if os.Getenv("STRIPE_SECRET_KEY") == "" {
		log.Fatal("❌ STRIPE_SECRET_KEY is not set in environment variables")
	}

	// Initialize PostgreSQL repository
	repo := repository.NewPaymentRepo()

	// Initialize RabbitMQ connection and exchange
	rabbitmq.InitRabbitMQ()

	// Initialize Stripe Gateway
	gateway.InitStripe() // ✅ your new stripe gateway init

	// Start Stripe Webhook Server
	webhook.StartStripeWebhookServer(repo) // ✅ handles payment_intent.succeeded, etc.

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50058")
	if err != nil {
		log.Fatalf("❌ Failed to listen on port 50058: %v", err)
	}

	grpcServer := grpc.NewServer()
	paymentpb.RegisterPaymentServiceServer(grpcServer, &handlers.PaymentService{Repo: repo})

	log.Println("💳 Payment Service (Stripe/COD) gRPC running on port :50058")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve gRPC: %v", err)
	}
}
