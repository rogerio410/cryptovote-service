package mongodb

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
)

type MongoDBCryptoRepository struct {
}

func NewMongoDBCryptoRepository() *MongoDBCryptoRepository {
	return &MongoDBCryptoRepository{}
}

func (m MongoDBCryptoRepository) GetAll(ctx context.Context) ([]domain.Cryptocurrency, error) {
	cryptos := []domain.Cryptocurrency{
		{Symbol: "BTC", Name: "Bitcoin"},
		{Symbol: "ETH", Name: "Ethereum"},
		{Symbol: "ADA", Name: "Cardano"},
	}

	return cryptos, nil
}
