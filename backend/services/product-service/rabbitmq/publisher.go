package rabbitmq

import (
	"encoding/json"
	"log"
	"product-service/models"
	"time"

	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func Init(channel *amqp.Channel) {
	ch = channel
}

func EmitLowStockEvent(p *models.Product) {
	if ch == nil {
		log.Println("❌ RabbitMQ channel not set")
		return
	}

	payload := map[string]interface{}{
		"product_id": p.ID,
		"quantity":   p.Quantity,
		"name":       p.Name,
		"brand":      p.Brand,
		"timestamp":  time.Now().Format(time.RFC3339),
	}

	body, _ := json.Marshal(payload)

	err := ch.Publish(
		"", "product.low_stock", false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("❌ Failed to publish low stock event: %v", err)
	} else {
		log.Printf("📤 Low stock event published for product %s", p.ID)
	}
}
