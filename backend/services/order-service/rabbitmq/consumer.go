package rabbitmq

import (
	"encoding/json"
	"order-service/repository"

	amqp "github.com/rabbitmq/amqp091-go"
)

func StartConsumer(repo *repository.OrderRepo) {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()

	ch.QueueDeclare("payment.success", true, false, false, false, nil)
	ch.QueueDeclare("payment.failed", true, false, false, false, nil)

	successMsgs, _ := ch.Consume("payment.success", "", true, false, false, false, nil)
	failedMsgs, _ := ch.Consume("payment.failed", "", true, false, false, false, nil)

	go func() {
		for msg := range successMsgs {
			var data map[string]string
			json.Unmarshal(msg.Body, &data)
			repo.UpdatePaymentStatus(data["order_id"], "success")
		}
	}()
	go func() {
		for msg := range failedMsgs {
			var data map[string]string
			json.Unmarshal(msg.Body, &data)
			repo.UpdatePaymentStatus(data["order_id"], "failed")
		}
	}()
}
