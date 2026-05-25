package logs

import (
	"context"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

func NewRepository() *Repository {

	collection := db.MongoClient.
		Database("fluxion_logs").
		Collection("logs")

	return &Repository{
		Collection: collection,
	}
}

func (r *Repository) Create(logData Log) error {

	_, err := r.Collection.InsertOne(
		context.Background(),
		logData,
	)
	return err
}

func (r *Repository) FindAll() ([]Log, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	var logs []Log
	for cursor.Next(context.Background()) {

		var logData Log
		err := cursor.Decode(&logData)
		if err != nil {
			return nil, err
		}
		logs = append(logs, logData)
	}

	return logs, nil
}