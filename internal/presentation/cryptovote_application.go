package presentation

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
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
	userRepo := mongodb.NewMongoDBUserRepository(mongoClient)

	// Concrete App
	application := app.Application{
		Queries: app.Queries{
			GetAllCrypto:      query.NewGetAllCriptoQuery(cryptoRepo),
			GetCryptoBySymbol: query.NewGetCriptoBySymbolQuery(cryptoRepo),
		},
		Commands: app.Commands{
			Vote:       command.NewVoteCommand(cryptoRepo, userRepo),
			RemoveVote: command.NewRemoveVoteCommand(cryptoRepo, userRepo),
		},
	}

	return application
}

func newMongoDBClient(ctx context.Context) *mongo.Client {
	envs, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	host := envs["MONGO_HOST"]
	port := envs["MONGO_PORT"]
	user := envs["MONGO_USER"]
	password := envs["MONGO_PASSWORD"]

	mongodbUri := ""
	if user != "" {
		mongodbUri = fmt.Sprintf("mongodb://%v:%v@%v:%v", user, password, host, port)
	} else {
		mongodbUri = fmt.Sprintf("mongodb://%v:%v", host, port)
	}

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
