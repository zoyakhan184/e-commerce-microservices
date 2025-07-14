package clients

import (
	imagepb "bff-service/proto/image"
	"log"
	"os"

	"google.golang.org/grpc"
)

var ImageClient imagepb.ImageServiceClient

func InitImageClient() {
	addr := os.Getenv("IMAGE_SERVICE_ADDR")
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC dial error (Image): %v", err)
	}
	ImageClient = imagepb.NewImageServiceClient(conn)
	log.Println("✅ Image client ready")
}
