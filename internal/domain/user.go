package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	username string             `bson:"username,omitempty" json:"username,omitempty"`
}
