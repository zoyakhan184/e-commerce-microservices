// rabbitmq/publisher.go
package rabbitmq

import (
	"encoding/json"
	"inventory-service/models"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func SetChannel(channel *amqp.Channel) {
	ch = channel
}

func EmitLowStockEvent(inv *models.Inventory) {
	if ch == nil {
		log.Println("⚠️ RabbitMQ channel is not set. Cannot emit low stock event.")
		return
	}

	payload := map[string]interface{}{
		"sku_id":     inv.SkuID,
		"product_id": inv.ProductID,
		"quantity":   inv.Quantity,
		"size":       inv.Size,
		"color":      inv.Color,
		"timestamp":  time.Now().Format(time.RFC3339),
	}

	body, _ := json.Marshal(payload)

	err := ch.Publish(
		"", "inventory.low_stock", false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("❌ Failed to emit low stock event: %v", err)
	} else {
		log.Printf("📤 Low stock event emitted for SKU: %s", inv.SkuID)
	}
}
