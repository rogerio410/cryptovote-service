package domain

type Cryptocurrency struct {
	Symbol string `bson:"symbol,omitempty"`
	Name   string `bson:"name,omitempty"`
	Votes  []Vote
}
