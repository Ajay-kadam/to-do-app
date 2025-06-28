package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title   string             `json:"title"`
	Details string             `json:"details"`
}
