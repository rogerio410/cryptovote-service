package presentation

import (
	"context"
	"fmt"
	"log"

	app "github.com/rogerio410/cryptovote-service/internal/application"
	"github.com/rogerio410/cryptovote-service/internal/application/command"
	"github.com/rogerio410/cryptovote-service/internal/application/query"
	"github.com/rogerio410/cryptovote-service/internal/infra/mongodb"
	"github.com/spf13/viper"
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
	host, _ := viper.Get("MONGO_HOST").(string)
	port, _ := viper.Get("MONGO_PORT").(string)
	user, _ := viper.Get("MONGO_USER").(string)
	password, _ := viper.Get("MONGO_PASSWORD").(string)

	mongodbUri := ""
	if user != "" {
		mongodbUri = fmt.Sprintf("mongodb+srv://%v:%v@%v/cryptovotes?retryWrites=true&w=majority&ssl=true&authSource=admin", user, password, host)
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
