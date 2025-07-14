package clients

import (
	cartpb "bff-service/proto/cart"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
)

var cartOnce sync.Once
var cartClient cartpb.CartServiceClient

func CartClient() cartpb.CartServiceClient {
	cartOnce.Do(func() {
		addr := os.Getenv("CART_SERVICE_ADDR")
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("gRPC dial error (Cart): %v", err)
		}
		cartClient = cartpb.NewCartServiceClient(conn)
		fmt.Println("✅ Cart client ready")
	})
	return cartClient
}
