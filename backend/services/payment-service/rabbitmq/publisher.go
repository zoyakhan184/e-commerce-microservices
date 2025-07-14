package rabbitmq

import (
    "encoding/json"
    "log"
    "os"

    "github.com/streadway/amqp"
)

var channel *amqp.Channel

func InitRabbitMQ() {
    conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
    if err != nil {
        panic(err)
    }

    ch, err := conn.Channel()
    if err != nil {
        panic(err)
    }

    channel = ch
}

func EmitPaymentEvent(eventType string, payload interface{}) {
    body, _ := json.Marshal(payload)
    err := channel.Publish(
        "", eventType, false, false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
    if err != nil {
        log.Printf("Failed to publish %s: %v", eventType, err)
    }
}
