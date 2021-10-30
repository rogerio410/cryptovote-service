package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	name string             `bson:"name,omitempty"`
}
