package rabbitmq

import (
	"encoding/json"
	"log"
	"order-service/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

var channel *amqp.Channel

func InitRabbitMQ() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()
	channel = ch
	ch.QueueDeclare("order.placed", true, false, false, false, nil)
}

func EmitOrderPlaced(order models.Order) {
	body, _ := json.Marshal(map[string]string{
		"order_id": order.ID,
		"user_id":  order.UserID,
	})
	channel.Publish("", "order.placed", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	log.Println("📤 Emitted order.placed event for order_id:", order.ID)
}
