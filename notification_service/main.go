package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

type UserRegisteredEvent struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func main() {
	log.Println("Starting Notification Service...")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumerGroup(
		[]string{"localhost:9092"},
		"notification-service-group",
		config,
	)
	if err != nil {
		log.Fatal("Failed to create consumer:", err)
	}
	defer consumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	topics := []string{"user-registered"}
	handler := ConsumerGroupHandler{}

	go func() {
		for {
			if err := consumer.Consume(ctx, topics, &handler); err != nil {
				log.Printf("Error: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	log.Println("Notification Service is running...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm

	log.Println("Shutting down...")
}

type ConsumerGroupHandler struct{}

func (h ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var event UserRegisteredEvent

		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Failed to parse message: %v", err)
			continue
		}

		log.Println("========================================")
		log.Printf("To: %s", event.Email)
		log.Printf("   Subject: Welcome!")
		log.Printf("   Body: Hello %s, you successfully registered!", event.Username)
		log.Println("========================================")

		session.MarkMessage(msg, "")
	}
	return nil
}
