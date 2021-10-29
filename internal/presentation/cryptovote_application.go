package presentation

import (
	"context"

	app "github.com/rogerio410/cryptovote-service/internal/application"
	"github.com/rogerio410/cryptovote-service/internal/application/query"
	"github.com/rogerio410/cryptovote-service/internal/infra/mongodb"
)

func NewCryptoApplication(ctx context.Context) app.Application {

	// create specific dependencies
	cryptoRepo := mongodb.NewMongoDBCryptoRepository()

	// Concrete App
	application := app.Application{
		Queries: app.Queries{
			AllCrypto: query.NewGetAllCriptoQuery(cryptoRepo),
		},
	}

	return application
}
