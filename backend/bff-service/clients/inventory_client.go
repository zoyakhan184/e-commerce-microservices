package clients

import (
	inventorypb "bff-service/proto/inventory"
	"log"
	"os"

	"google.golang.org/grpc"
)

var InventoryClient inventorypb.InventoryServiceClient

func InitInventoryClient() {
	addr := os.Getenv("INVENTORY_SERVICE_ADDR")
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC dial error (Inventory): %v", err)
	}
	InventoryClient = inventorypb.NewInventoryServiceClient(conn)
	log.Println("✅ Inventory client ready")
}
