package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

var channel *amqp.Channel

// InitRabbitMQ initializes the connection and channel
func InitRabbitMQ(connStr string) {
	conn, err := amqp.Dial(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	channel = ch
	log.Println("📡 RabbitMQ connected and channel opened")
}

// Publish sends a message to the specified queue
func Publish(event string, payload map[string]interface{}) {
	if channel == nil {
		log.Println("❌ RabbitMQ channel not initialized")
		return
	}

	// Ensure the queue exists (declares it if not)
	_, err := channel.QueueDeclare(
		event, // name
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Printf("❌ Queue declare failed: %v", err)
		return
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("❌ JSON marshal failed: %v", err)
		return
	}

	err = channel.Publish(
		"",    // exchange
		event, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("❌ Failed to publish message: %v", err)
	} else {
		log.Printf("✅ Published event '%s': %s", event, string(body))
	}
}
