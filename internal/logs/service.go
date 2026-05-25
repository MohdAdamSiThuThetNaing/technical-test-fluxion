package logs

import (
	"context"

	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
)

type Service struct{}

func (s *Service) GetLogs() ([]Log, error) {

	collection := db.MongoClient.
		Database("fluxion_logs").
		Collection("logs")

	cursor, err := collection.Find(
		context.Background(),
		map[string]interface{}{},
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