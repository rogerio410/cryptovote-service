package mongodb

import (
	"context"

	"github.com/rogerio410/cryptovote-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBUserRepository struct {
	userCollection *mongo.Collection
}

func NewMongoDBUserRepository(mongoClient *mongo.Client) *MongoDBUserRepository {
	database := mongoClient.Database("cryptovote")

	usersCollection := database.Collection("users")

	return &MongoDBUserRepository{
		userCollection: usersCollection,
	}
}

func (m MongoDBUserRepository) GetByName(ctx context.Context, name string) (domain.User, error) {

	var user domain.User

	result := m.userCollection.FindOne(ctx, bson.M{"username": name})

	err := result.Decode(&user)

	return user, err
}
