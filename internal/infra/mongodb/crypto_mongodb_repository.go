package mongodb

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBCryptoRepository struct {
	cryptoCollection *mongo.Collection
	voteCollection   *mongo.Collection
}

func NewMongoDBCryptoRepository(mongoClient *mongo.Client) *MongoDBCryptoRepository {
	database := mongoClient.Database("cryptovote")

	cryptoCollection := database.Collection("cryptos")
	voteCollection := database.Collection("votes")

	return &MongoDBCryptoRepository{
		cryptoCollection: cryptoCollection,
		voteCollection:   voteCollection,
	}
}

func (m MongoDBCryptoRepository) GetAll(ctx context.Context) ([]domain.Cryptocurrency, error) {

	cursor, err := m.cryptoCollection.Find(
		ctx, bson.D{},
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

	var crypto domain.Cryptocurrency

	result := m.cryptoCollection.FindOne(ctx, bson.M{"symbol": symbol})

	err := result.Decode(&crypto)

	return crypto, err
}

func (m MongoDBCryptoRepository) GetVoteByCryptoAndUser(ctx context.Context, crypto domain.Cryptocurrency, user domain.User) (domain.Vote, error) {
	filter := bson.M{
		"crypto_id": bson.M{"$eq": crypto.ID},
		"user_id":   bson.M{"$eq": user.ID},
	}

	result := m.voteCollection.FindOne(ctx, filter)

	var vote domain.Vote

	err := result.Decode(&vote)

	return vote, err
}

func (m MongoDBCryptoRepository) AddVote(ctx context.Context, vote domain.Vote) error {
	filter := bson.M{
		"crypto_id": bson.M{"$eq": vote.CryptoID},
		"user_id":   bson.M{"$eq": vote.UserID},
	}
	data := bson.M{
		"$set": bson.M{
			"crypto_id": vote.CryptoID,
			"user_id":   vote.UserID,
			"vote":      vote.Vote,
		},
	}

	upsert := true
	result, err := m.voteCollection.UpdateOne(
		ctx,
		filter,
		data,
		&options.UpdateOptions{Upsert: &upsert},
	)

	if err != nil {
		return err
	}

	// Update votes count into
	if result.MatchedCount == 0 { // New Vote
		if vote.IsUpVote() {
			m.incUpVote(ctx, vote, 1)
		} else {
			m.incUpVote(ctx, vote, -1)
		}
	} else {
		if result.ModifiedCount == 1 {
			if vote.IsUpVote() {
				m.incUpVote(ctx, vote, 2)
			} else {
				m.incUpVote(ctx, vote, -2)
			}
		}
	}

	return nil
}

func (m MongoDBCryptoRepository) incUpVote(ctx context.Context, vote domain.Vote, value int32) {
	filter := bson.M{
		"_id": bson.M{"$eq": vote.CryptoID},
	}
	data := bson.M{
		"$inc": bson.M{
			"votes": value,
		},
	}

	m.cryptoCollection.UpdateOne(
		ctx,
		filter,
		data,
	)

}

func (m MongoDBCryptoRepository) RemoveVote(ctx context.Context, vote domain.Vote) error {
	filter := bson.M{
		"_id": vote.ID,
	}
	result, err := m.voteCollection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 1 {
		if vote.IsUpVote() {
			m.incUpVote(ctx, vote, -1)
		} else {
			m.incUpVote(ctx, vote, 1)
		}
	}

	return nil
}
