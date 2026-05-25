package queue

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/logs"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume() {

	conn, err := amqp.Dial(
		os.Getenv("RABBITMQ_URL"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	_, err = ch.QueueDeclare(
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
		"user_logs",
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

	collection := db.MongoClient.
		Database("fluxion_logs").
		Collection("logs")

	for msg := range msgs {

		var logData logs.Log

		err := json.Unmarshal(
			msg.Body,
			&logData,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		_, err = collection.InsertOne(
			context.Background(),
			logData,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Log inserted into MongoDB")
	}
}