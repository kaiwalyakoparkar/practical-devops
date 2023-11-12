package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Describes how the data will be stored in the database
type Netflix struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty" bson:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty" bson:"watched,omitempty"`
}
