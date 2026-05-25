package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(
			os.Getenv("MONGO_URI"),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client

	log.Println("MongoDB connected")
}