package queue

import (
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(data interface{}) {

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Printf("rabbitmq connection error: %v", err)
		return
	}

	defer conn.Close()
	ch, err := conn.Channel()

	if err != nil {
		log.Printf("rabbitmq channel error: %v", err)
		return
	}

	defer ch.Close()

	_, err = ch.QueueDeclare(
		"user_logs",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("queue declare error: %v", err)
		return
	}

	body, err := json.Marshal(data)

	if err != nil {
		log.Printf("json marshal error: %v", err)
		return
	}

	err = ch.Publish(
		"",
		"user_logs",
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			Body:         body,
		},
	)

	if err != nil {
		log.Printf("message publish error: %v", err)
		return
	}

	log.Println("event published successfully")
}