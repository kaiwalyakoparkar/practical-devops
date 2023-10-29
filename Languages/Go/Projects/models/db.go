package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Describes how the data will be stored in the database
type Netflix struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Movie   string             `bson:"movie,omitempty"`
	Watched bool               `bson:"watched,omitempty"`
}
