package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/jamsxd/marvel-api/internal/marvel/character/application"
	"github.com/jamsxd/marvel-api/internal/marvel/character/domain"
	"github.com/jamsxd/marvel-api/internal/marvel/character/infrastructure/persistence"
	"github.com/jamsxd/marvel-api/internal/marvel/character/infrastructure/transport"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	var (
		port   = os.Getenv("PORT")
		dbName = os.Getenv("DB_NAME")
		dbHost = os.Getenv("DB_HOST")
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conexion := "mongodb://" + dbUser + ":" + dbPass + "@" + dbHost + "/"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conexion))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	var (
		repo = persistence.NewMongoRepository(client.Database(dbName))
		svc  = domain.NewBasicCharacterService(repo)
		end  = application.NewBasicCharacterEndpoint(svc)
		s    = transport.NewServer(*end)
	)

	http.ListenAndServe(":"+port, s)

}
