package mongodb

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBCryptoRepository struct {
	mongoClient *mongo.Client
}

func NewMongoDBCryptoRepository(mongoClient *mongo.Client) *MongoDBCryptoRepository {
	return &MongoDBCryptoRepository{mongoClient: mongoClient}
}

func (m MongoDBCryptoRepository) GetAll(ctx context.Context) ([]domain.Cryptocurrency, error) {

	cryptoCollection := m.mongoClient.Database("cryptovote").Collection("cryptos")

	// projection := bson.D{
	// 	// {Key: "votes", Value: 0},
	// 	{Key: "name", Value: 1},
	// 	{Key: "symbol", Value: 1},
	// }
	cursor, err := cryptoCollection.Find(
		ctx, bson.D{},
		// options.Find().SetProjection(projection),
	)

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

func (m MongoDBCryptoRepository) GetVoteByCryptoAndUser(ctx context.Context, crypto domain.Cryptocurrency, user domain.User) (domain.Vote, error) {
	voteCollection := m.mongoClient.Database("cryptovote").Collection("votes")

	filter := bson.M{
		"crypto": bson.M{"$eq": crypto.ID},
		"user":   bson.M{"$eq": user.ID},
	}

	result := voteCollection.FindOne(ctx, filter)

	var vote domain.Vote

	err := result.Decode(&vote)

	return vote, err
}

func (m MongoDBCryptoRepository) AddVote(ctx context.Context, vote domain.Vote) error {
	voteCollection := m.mongoClient.Database("cryptovote").Collection("votes")

	filter := bson.M{
		"crypto": bson.M{"$eq": vote.CryptoID},
		"user":   bson.M{"$eq": vote.UserID},
	}
	data := bson.M{
		"$set": bson.M{
			"crypto": vote.CryptoID,
			"user":   vote.UserID,
			"vote":   vote.Vote,
		},
	}

	upsert := true
	_, err := voteCollection.UpdateOne(
		ctx,
		filter,
		data,
		&options.UpdateOptions{Upsert: &upsert},
	)

	if err != nil {
		return err
	}

	return nil
}

func (m MongoDBCryptoRepository) RemoveVote(ctx context.Context, vote domain.Vote) error {
	voteCollection := m.mongoClient.Database("cryptovote").Collection("votes")

	filter := bson.M{
		"_id": vote.ID,
	}
	_, err := voteCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
