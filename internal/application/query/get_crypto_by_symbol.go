package query

import (
	"context"
	"errors"

	"github.com/rogerio410/cryptovote-service/internal/domain"
	"github.com/rogerio410/cryptovote-service/internal/domain/repository"
)

type GetCriptoBySymbolQuery struct {
	cryptoRepository repository.CryptoRepository
}

func NewGetCriptoBySymbolQuery(cryptoRepository repository.CryptoRepository) GetCriptoBySymbolQuery {
	if cryptoRepository == nil {
		panic("Must provide a Cryptocurrency Repository")
	}

	return GetCriptoBySymbolQuery{cryptoRepository}
}

func (g GetCriptoBySymbolQuery) Execute(ctx context.Context, symbol string) (domain.Cryptocurrency, error) {
	crypto, err := g.cryptoRepository.GetBySymbol(ctx, symbol)
	if err != nil {
		return domain.Cryptocurrency{}, errors.New("Invalid Cryptocurrency symbol!")
	}
	return crypto, nil
}
