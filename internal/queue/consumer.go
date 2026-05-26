package queue

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume() {

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

	msgs, err := ch.Consume(
		"user_logs",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Printf("queue consume error: %v", err)
		return
	}

	collection := db.MongoClient.Database("fluxion_logs").Collection("logs")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.Println("RabbitMQ consumer started...")

	for {

		select {

		case sig := <-sigChan:
			log.Printf("shutdown signal received: %v", sig)
			return

		case msg, ok := <-msgs:

			if !ok {
				log.Println("message channel closed")
				return
			}

			var logData logs.Log
			err := json.Unmarshal(msg.Body, &logData)

			if err != nil {
				log.Printf("json unmarshal error: %v", err)
				msg.Nack(false, false)
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			_, err = collection.InsertOne(ctx, logData)
			cancel()

			if err != nil {
				log.Printf("mongodb insert error: %v", err)
				msg.Nack(false, true)
				continue
			}

			err = msg.Ack(false)
			if err != nil {
				log.Printf("message ack error: %v", err)
				continue
			}

			log.Println("log inserted into MongoDB")
		}
	}
}