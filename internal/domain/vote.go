package domain

type Vote struct {
	User User `bson:"user,omitempty"`
	Vote bool `bson:"vote,omitempty"`
}
