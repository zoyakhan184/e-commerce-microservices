package clients

import (
	reviewpb "bff-service/proto/review"
	"log"
	"os"

	"google.golang.org/grpc"
)

var ReviewClient reviewpb.ReviewServiceClient

func InitReviewClient() {
	addr := os.Getenv("REVIEW_SERVICE_ADDR")
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC dial error (Review): %v", err)
	}
	ReviewClient = reviewpb.NewReviewServiceClient(conn)
	log.Println("✅ Review client ready")
}
