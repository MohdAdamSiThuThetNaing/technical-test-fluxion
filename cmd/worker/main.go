package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	db.ConnectMongo()

	log.Println("MongoDB connected")

	time.Sleep(10 * time.Second)

	log.Println("Connecting RabbitMQ...")

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	log.Println("RabbitMQ connected")

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"user_logs",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	collection := db.MongoClient.Database("fluxion_logs").Collection("logs")

	log.Println("Worker started...")

	forever := make(chan bool)

	go func() {

		for msg := range msgs {

			var logData logs.Log

			err := json.Unmarshal(msg.Body, &logData)

			if err != nil {
				log.Println(err)
				continue
			}

			_, err = collection.InsertOne(context.Background(), logData)

			if err != nil {
				log.Println(err)
				continue
			}

			log.Println("Log saved:", logData.Event)
		}
	}()

	<-forever
}