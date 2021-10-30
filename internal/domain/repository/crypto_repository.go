package repository

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
)

type CryptoRepository interface {
	GetAll(ctx context.Context) ([]domain.Cryptocurrency, error)
	GetBySymbol(ctx context.Context, symbol string) (domain.Cryptocurrency, error)
	GetVoteByCryptoAndUser(ctx context.Context, crypto domain.Cryptocurrency, user domain.User) (domain.Vote, error)
	AddVote(ctx context.Context, vote domain.Vote) error
	RemoveVote(ctx context.Context, vote domain.Vote) error
}
