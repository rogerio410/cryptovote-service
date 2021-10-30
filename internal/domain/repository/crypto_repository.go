package repository

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
)

type CryptoRepository interface {
	GetAll(ctx context.Context) ([]domain.Cryptocurrency, error)
	GetBySymbol(ctx context.Context, symbol string) (domain.Cryptocurrency, error)
	AddVote(ctx context.Context, symbol string, vote domain.Vote) error
	RemoveVote(ctx context.Context, symbol string, user domain.User) error
}
