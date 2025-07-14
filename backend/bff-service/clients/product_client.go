package clients

import (
	productpb "bff-service/proto/product"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
)

var productOnce sync.Once
var productClient productpb.ProductServiceClient

func ProductClient() productpb.ProductServiceClient {
	productOnce.Do(func() {
		addr := os.Getenv("PRODUCT_SERVICE_ADDR")
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("gRPC dial error (Product): %v", err)
		}
		productClient = productpb.NewProductServiceClient(conn)
		fmt.Println("✅ Product client ready")
	})
	return productClient
}
