package repository

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
)

type CryptocurrencyRepository interface {
	GetAll(ctx context.Context) ([]domain.Cryptocurrency, error)
	GetOne(ctx context.Context, symbol string) (*domain.Cryptocurrency, error)
	AddVote(ctx context.Context, vote domain.Vote) error
	UpdateVote(ctx context.Context, voteID string, newValue bool) error
	RemoveVote(ctx context.Context, voteID string) error
}
