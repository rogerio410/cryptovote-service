package command

import (
	"context"
	"errors"

	"github.com/rogerio410/cryptovote-service/internal/domain/repository"
)

type RemoveVoteCommand struct {
	cryptoRepository repository.CryptoRepository
	userRepository   repository.UserRepository
}

func NewRemoveVoteCommand(cryptoRepository repository.CryptoRepository, userRepository repository.UserRepository) RemoveVoteCommand {
	if cryptoRepository == nil {
		panic("Must provide a Cryptocurrency Repository")
	}

	return RemoveVoteCommand{cryptoRepository, userRepository}
}

func (u RemoveVoteCommand) Execute(ctx context.Context, symbol string, username string) error {
	// TODO: see business logic
	crypto, err := u.cryptoRepository.GetBySymbol(ctx, symbol)

	if err != nil {
		return errors.New("Invalid Cryptocurrency symbol!")
	}

	user, user_err := u.userRepository.GetByName(ctx, username)

	if user_err != nil {
		return errors.New("Invalid username!")
	}

	vote, err := u.cryptoRepository.GetVoteByCryptoAndUser(ctx, crypto, user)

	if err != nil {
		return errors.New("Vote not found!")
	}

	add_err := u.cryptoRepository.RemoveVote(ctx, vote)

	if add_err != nil {
		return errors.New("Error on remove vote!")
	}

	return nil
}
