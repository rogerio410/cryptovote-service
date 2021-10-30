package command

import (
	"context"
	"errors"

	"github.com/rogerio410/cryptovote-service/internal/domain"
	"github.com/rogerio410/cryptovote-service/internal/domain/repository"
)

type VoteCommand struct {
	cryptoRepository repository.CryptoRepository
	userRepository   repository.UserRepository
}

func NewVoteCommand(cryptoRepository repository.CryptoRepository, userRepository repository.UserRepository) VoteCommand {
	if cryptoRepository == nil {
		panic("Must provide a Cryptocurrency Repository")
	}

	return VoteCommand{cryptoRepository, userRepository}
}

func (u VoteCommand) Execute(ctx context.Context, symbol string, value string, username string) error {
	crypto, err := u.cryptoRepository.GetBySymbol(ctx, symbol)
	if err != nil {
		return errors.New("Invalid Cryptocurrency symbol!")
	}

	user, user_err := u.userRepository.GetByName(ctx, username)
	if user_err != nil {
		return errors.New("Invalid username!")
	}

	vote := domain.Vote{Vote: value, UserID: user.ID, CryptoID: crypto.ID}

	// add new one
	add_err := u.cryptoRepository.AddVote(ctx, vote)
	if add_err != nil {
		return errors.New("Error on add new vote!")
	}

	return nil
}
