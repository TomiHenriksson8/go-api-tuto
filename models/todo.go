package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
	UserID    primitive.ObjectID `json:"userID,omitempty" bson:"userID,omitempty"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type TodoCompletionRequest struct {
	Completed bool `json:"completed"`
}
