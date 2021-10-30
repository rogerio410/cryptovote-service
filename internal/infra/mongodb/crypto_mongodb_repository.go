package mongodb

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBCryptoRepository struct {
	mongoClient *mongo.Client
}

func NewMongoDBCryptoRepository(mongoClient *mongo.Client) *MongoDBCryptoRepository {
	return &MongoDBCryptoRepository{mongoClient: mongoClient}
}

func (m MongoDBCryptoRepository) GetAll(ctx context.Context) ([]domain.Cryptocurrency, error) {

	cryptoCollection := m.mongoClient.Database("cryptovote").Collection("cryptos")

	cursor, err := cryptoCollection.Find(ctx, bson.D{})

	if err != nil {
		panic(err)
	}
	var cryptos []domain.Cryptocurrency

	if err = cursor.All(ctx, &cryptos); err != nil {
		panic(err)
	}

	return cryptos, nil
}

func (m MongoDBCryptoRepository) GetBySymbol(ctx context.Context, symbol string) (domain.Cryptocurrency, error) {

	cryptoCollection := m.mongoClient.Database("cryptovote").Collection("cryptos")

	var crypto domain.Cryptocurrency

	result := cryptoCollection.FindOne(ctx, bson.M{"symbol": symbol})

	err := result.Decode(&crypto)

	return crypto, err
}

func (m MongoDBCryptoRepository) AddVote(ctx context.Context, symbol string, vote domain.Vote) error {
	cryptoCollection := m.mongoClient.Database("cryptovote").Collection("cryptos")

	filter := bson.M{"symbol": symbol}
	update := bson.M{"$push": bson.M{"votes": vote}}
	_, err := cryptoCollection.UpdateOne(
		ctx,
		filter,
		update,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m MongoDBCryptoRepository) RemoveVote(ctx context.Context, symbol string, user domain.User) error {
	cryptoCollection := m.mongoClient.Database("cryptovote").Collection("cryptos")

	filter := bson.M{"symbol": symbol}
	update := bson.M{"$pull": bson.M{"votes": bson.M{"user": user.ID}}}
	_, err := cryptoCollection.UpdateOne(
		ctx,
		filter,
		update,
	)

	if err != nil {
		return err
	}

	return nil
}
