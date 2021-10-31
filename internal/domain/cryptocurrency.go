package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cryptocurrency struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Symbol string             `bson:"symbol,omitempty" json:"symbol,omitempty`
	Name   string             `bson:"name,omitempty" json:"name,omitempty"`
	Votes  int32              `bson:"votes,omitempty" json:"votes,omitempty"`
}
