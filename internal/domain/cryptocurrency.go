package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cryptocurrency struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Symbol string             `bson:"symbol,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Votes  []Vote             `bson:"votes,omitempty"`
}
