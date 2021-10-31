package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vote struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	CryptoID primitive.ObjectID `bson:"crypto_id,omitempty" json:"crypto_id,omitempty"`
	Vote     string             `bson:"vote,omitempty" json:"vote,omitempty"`
}

func (v Vote) IsUpVote() bool {
	return v.Vote == "UP"
}

func (v Vote) IsDownVote() bool {
	return !v.IsUpVote()
}
