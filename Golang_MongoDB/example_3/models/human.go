package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Human struct {
	ID       	primitive.ObjectID		`json:"id" bson:"_id,omitempty"`
	Firstname   string            		`json:"firstname" bson:"firstname,omitempty"`
	Lastname	string             		`json:"lastname" bson:"lastname,omitempty"`
	Age 		int64            		`json:"age" bson:"age,omitempty"`
}
