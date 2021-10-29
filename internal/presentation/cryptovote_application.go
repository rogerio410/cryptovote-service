package presentation

import (
	"context"
	"fmt"
	"log"

	app "github.com/rogerio410/cryptovote-service/internal/application"
	"github.com/rogerio410/cryptovote-service/internal/application/command"
	"github.com/rogerio410/cryptovote-service/internal/application/query"
	"github.com/rogerio410/cryptovote-service/internal/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewCryptoApplication(ctx context.Context) app.Application {

	// create specific dependencies

	mongoClient := newMongoDBClient(ctx)
	cryptoRepo := mongodb.NewMongoDBCryptoRepository(mongoClient)

	// Concrete App
	application := app.Application{
		Queries: app.Queries{
			AllCrypto: query.NewGetAllCriptoQuery(cryptoRepo),
		},
		Commands: app.Commands{
			Vote: command.NewVoteCommand(cryptoRepo),
		},
	}

	return application
}

func newMongoDBClient(ctx context.Context) *mongo.Client {

	mongodbUri := fmt.Sprintf("mongodb://localhost:27017/?readPreference=primary&directConnection=true&ssl=false")
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongodbUri))
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return mongoClient
}
