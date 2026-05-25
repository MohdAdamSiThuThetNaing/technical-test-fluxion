package queue

import (
	"encoding/json"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(data interface{}) {

	conn, err := amqp.Dial(
		os.Getenv("RABBITMQ_URL"),
	)

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return
	}

	body, _ := json.Marshal(data)

	err = ch.Publish(
		"",
		"user_logs",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Event published")
}