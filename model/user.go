package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"Id,omitempty"`
	FirstName string             `json:"FirstName"`
	LastName  string             `json:"LastName"`
}
