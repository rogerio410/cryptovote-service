package query

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
	"github.com/rogerio410/cryptovote-service/internal/domain/repository"
)

type GetAllCriptoQuery struct {
	cryptoRepository repository.CryptoRepository
}

func NewGetAllCriptoQuery(cryptoRepository repository.CryptoRepository) GetAllCriptoQuery {
	if cryptoRepository == nil {
		panic("Must provide a Cryptocurrency Repository")
	}

	return GetAllCriptoQuery{cryptoRepository}
}

func (g GetAllCriptoQuery) Execute(ctx context.Context) ([]domain.Cryptocurrency, error) {
	return g.cryptoRepository.GetAll(ctx)
}
