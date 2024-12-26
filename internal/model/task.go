package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task string             `json:"task,omitempty"`
	Done bool               `json:"done,omitempty"`
}
