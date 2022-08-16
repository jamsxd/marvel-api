package persistence

import (
	"context"

	"github.com/jamsxd/marvel-api/internal/marvel/character/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) domain.Repository {

	return &MongoRepository{
		db: db.Collection("characters"),
	}
}

func (r *MongoRepository) Save(ctx context.Context, character domain.Character) (*domain.Character, error) {
	_, err := r.db.InsertOne(ctx, character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}
