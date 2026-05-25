package logs

import "time"

type Log struct {
	Event       string    `json:"event" bson:"event"`
	UserID      string    `json:"user_id" bson:"user_id"`
	UserEmail   string    `json:"user_email" bson:"user_email"`
	UpdatedName string    `json:"updated_name" bson:"updated_name"`
	ActionBy    string    `json:"action_by" bson:"action_by"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}