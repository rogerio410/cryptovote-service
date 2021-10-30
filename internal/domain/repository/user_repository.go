package repository

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
)

type UserRepository interface {
	GetByName(ctx context.Context, name string) (domain.User, error)
}
