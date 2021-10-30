package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vote struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	User primitive.ObjectID `bson:"user,omitempty"`
	Vote string             `bson:"vote,omitempty"`
}
