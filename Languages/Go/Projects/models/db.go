package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Describes how the data will be stored in the database
type Netflix struct {
	ID      primitive.ObjectID `json: "_id,omitempty" bson: "_id, omitempty" `
	Movie   string             `json: "movie, omitempty"`
	Watched bool               `json: "watched, omitempty"`
}
