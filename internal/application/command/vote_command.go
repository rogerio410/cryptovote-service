package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/rogerio410/cryptovote-service/internal/domain/repository"
)

type VoteCommand struct {
	cryptoRepository repository.CryptocurrencyRepository
}

func NewVoteCommand(cryptoRepository repository.CryptocurrencyRepository) VoteCommand {
	if cryptoRepository == nil {
		panic("Must provide a Cryptocurrency Repository")
	}

	return VoteCommand{cryptoRepository}
}

func (u VoteCommand) Execute(ctx context.Context, symbol string, value string) error {
	// TODO: see business logic
	crypto, err := u.cryptoRepository.GetCurrencyBySymbol(ctx, symbol)

	if err != nil {
		return errors.New("Invalid Cryptocurrency symbol!")
	}

	fmt.Println(crypto.Name)

	return err
}
