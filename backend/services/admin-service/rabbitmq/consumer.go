package rabbitmq

import (
	"admin-service/models"
	"admin-service/repository"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func StartListeners(repo *repository.AdminRepo) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("RabbitMQ connect error: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("RabbitMQ channel error: %v", err)
	}

	listen(ch, "user.registered", func(body []byte) {
		var payload struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}
		if err := json.Unmarshal(body, &payload); err == nil {
			msg := fmt.Sprintf("User %s registered", payload.Email)
			repo.AddActivity(models.ActivityLog{
				Type:      "user",
				Message:   msg,
				Timestamp: time.Now().Format(time.RFC3339),
			})
			log.Println("📥 user.registered ->", msg)
		}
	})

	listen(ch, "order.placed", func(body []byte) {
		var payload struct {
			OrderID     string  `json:"order_id"`
			TotalAmount float64 `json:"total_amount"`
		}
		if err := json.Unmarshal(body, &payload); err == nil {
			msg := fmt.Sprintf("Order #%s placed ($%.2f)", payload.OrderID, payload.TotalAmount)
			repo.AddActivity(models.ActivityLog{
				Type:      "order",
				Message:   msg,
				Timestamp: time.Now().Format(time.RFC3339),
			})
			log.Println("📥 order.placed ->", msg)
		}
	})

	listen(ch, "payment.success", func(body []byte) {
		var payload struct {
			OrderID string  `json:"order_id"`
			Amount  float64 `json:"amount"`
		}
		if err := json.Unmarshal(body, &payload); err == nil {
			msg := fmt.Sprintf("Payment received for Order #%s ($%.2f)", payload.OrderID, payload.Amount)
			repo.AddActivity(models.ActivityLog{
				Type:      "payment",
				Message:   msg,
				Timestamp: time.Now().Format(time.RFC3339),
			})
			log.Println("📥 payment.success ->", msg)
		}
	})

	listen(ch, "inventory.low_stock", func(body []byte) {
		var payload struct {
			SkuID     string `json:"sku_id"`
			ProductID string `json:"product_id"`
			Quantity  int    `json:"quantity"`
		}
		if err := json.Unmarshal(body, &payload); err == nil {
			msg := fmt.Sprintf("Low stock alert: SKU %s - %d left", payload.SkuID, payload.Quantity)
			repo.AddActivity(models.ActivityLog{
				Type:      "inventory",
				Message:   msg,
				Timestamp: time.Now().Format(time.RFC3339),
			})
			log.Println("📥 inventory.low_stock ->", msg)
		}
	})
}

// Generic listener
func listen(ch *amqp.Channel, queue string, handler func([]byte)) {
	_, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue %s: %v", queue, err)
	}

	msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to consume queue %s: %v", queue, err)
	}

	go func() {
		for msg := range msgs {
			handler(msg.Body)
		}
	}()
}
