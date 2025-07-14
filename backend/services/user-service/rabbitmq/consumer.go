package consumer

import (
	"log"
	"os"
	"user-service/models"
	"user-service/repository"

	"github.com/streadway/amqp"
)

func StartUserEventListener(repo repository.UserRepository) {
	go func() {
		conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
		if err != nil {
			log.Fatalf("RabbitMQ connection failed: %v", err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("Channel open failed: %v", err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare("user.registered", true, false, false, false, nil)
		if err != nil {
			log.Fatalf("Queue declare failed: %v", err)
		}

		msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
		if err != nil {
			log.Fatalf("Consumer registration failed: %v", err)
		}

		log.Println("📥 Listening for `user.registered` events...")
		for msg := range msgs {
			userID := string(msg.Body)
			log.Printf("Received user.registered for user_id=%s", userID)

			profile := &models.Profile{
				UserID: userID,
			}

			if err := repo.CreateProfile(profile); err != nil {
				log.Printf("❌ Failed to create profile: %v", err)
			} else {
				log.Printf("✅ Profile created for user_id=%s", userID)
			}
		}
	}()
}
