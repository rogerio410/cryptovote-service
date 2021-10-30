package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vote struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty"`
	CryptoID primitive.ObjectID `bson:"crypto_id,omitempty"`
	Vote     string             `bson:"vote,omitempty"`
}
