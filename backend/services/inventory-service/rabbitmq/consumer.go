package rabbitmq

import (
	"encoding/json"
	"fmt"
	"inventory-service/repository"
	"os"

	"github.com/streadway/amqp"
)

var repo repository.InventoryRepository

func InitRabbitMQ(inventoryRepo repository.InventoryRepository) {
	repo = inventoryRepo

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		fmt.Printf("❌ Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("❌ Failed to open RabbitMQ channel: %v", err)
	}

	// Pass channel to publisher
	SetChannel(ch)

	// Declare queues
	_, _ = ch.QueueDeclare("order.placed", true, false, false, false, nil)
	_, _ = ch.QueueDeclare("inventory.low_stock", true, false, false, false, nil)

	msgs, err := ch.Consume("order.placed", "", true, false, false, false, nil)
	if err != nil {
		fmt.Printf("❌ Failed to consume order.placed: %v", err)
	}

	go func() {
		for d := range msgs {
			var payload struct {
				SkuID          string `json:"sku_id"`
				QuantityChange int    `json:"quantity"`
			}

			if err := json.Unmarshal(d.Body, &payload); err != nil {
				fmt.Printf("❌ Failed to parse message: %v", err)
				continue
			}

			// Update stock
			if err := repo.UpdateStock(payload.SkuID, -payload.QuantityChange); err != nil {
				fmt.Printf("❌ Stock update failed: %v", err)
				continue
			}
			fmt.Printf("✅ Stock updated for SKU: %s", payload.SkuID)

			// Check for low stock and emit event
			inv, err := repo.GetBySKU(payload.SkuID)
			if err == nil && inv.Quantity <= 5 {
				fmt.Printf("⚠️ Low stock detected for SKU: %s", inv.SkuID)
				EmitLowStockEvent(inv)
			}
		}
	}()
}
