package clients

import (
	orderpb "bff-service/proto/order"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
)

var orderOnce sync.Once
var orderClient orderpb.OrderServiceClient

func OrderClient() orderpb.OrderServiceClient {
	orderOnce.Do(func() {
		addr := os.Getenv("ORDER_SERVICE_ADDR")
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("gRPC dial error (Order): %v", err)
		}
		orderClient = orderpb.NewOrderServiceClient(conn)
		fmt.Println("✅ Order client ready")
	})
	return orderClient
}
