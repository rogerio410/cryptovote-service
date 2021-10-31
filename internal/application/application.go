package application

import (
	"github.com/rogerio410/cryptovote-service/internal/application/command"
	"github.com/rogerio410/cryptovote-service/internal/application/query"
)

type Application struct {
	Queries  Queries
	Commands Commands
}

type Queries struct {
	GetAllCrypto      query.GetAllCriptoQuery
	GetCryptoBySymbol query.GetCriptoBySymbolQuery
}

type Commands struct {
	Vote       command.VoteCommand
	RemoveVote command.RemoveVoteCommand
}
