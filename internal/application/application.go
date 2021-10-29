package application

import (
	"github.com/rogerio410/cryptovote-service/internal/application/query"
)

type Application struct {
	Queries Queries
}

type Queries struct {
	AllCrypto query.GetAllCriptoQuery
}
