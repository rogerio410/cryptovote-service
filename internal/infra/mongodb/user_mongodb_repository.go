package mongodb

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBUserRepository struct {
	mongoClient *mongo.Client
}

func NewMongoDBUserRepository(mongoClient *mongo.Client) *MongoDBCryptoRepository {
	return &MongoDBCryptoRepository{mongoClient: mongoClient}
}

func (m MongoDBCryptoRepository) GetByName(ctx context.Context, name string) (domain.User, error) {

	usersCollection := m.mongoClient.Database("cryptovote").Collection("users")

	var user domain.User

	result := usersCollection.FindOne(ctx, bson.M{"name": name})

	err := result.Decode(&user)

	return user, err
}
