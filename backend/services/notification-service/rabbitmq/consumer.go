package rabbitmq

import (
	"encoding/json"
	"log"
	"notification-service/config"
	"notification-service/repository"
	"os"

	"github.com/streadway/amqp"
)

func StartRabbitMQConsumer(repo *repository.NotificationRepo) {
	conn, _ := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	ch, _ := conn.Channel()

	q, _ := ch.QueueDeclare("notification-queue", true, false, false, false, nil)
	_ = ch.QueueBind(q.Name, "#", "amq.topic", false, nil)

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	go func() {
		for d := range msgs {
			log.Printf("🔔 Event Received: %s", d.RoutingKey)
			var data map[string]string
			_ = json.Unmarshal(d.Body, &data)

			to := data["email"]
			subject := "Notification"
			body := "You have a new event: " + d.RoutingKey

			if to != "" {
				_ = config.SendEmail(to, subject, body)
			}
		}
	}()
}
