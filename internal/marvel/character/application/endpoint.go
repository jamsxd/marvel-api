package application

import (
	"context"

	"github.com/jamsxd/marvel-api/internal/marvel/character/domain"
)

type Endpoint func(ctx context.Context, request interface{}) (interface{}, error)

type CharacterEndpoint struct {
	CreateCharacter Endpoint
}

func NewBasicCharacterEndpoint(svc domain.CharacterService) *CharacterEndpoint {
	return &CharacterEndpoint{
		CreateCharacter: makeCreateCharacter(svc),
	}
}

type CreateCharacterRequest struct {
	Character domain.Character
}

type CreateCharacterResponse struct {
	Character *domain.Character
}

func makeCreateCharacter(svc domain.CharacterService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCharacterRequest)
		character, err := svc.Create(ctx, req.Character)
		return CreateCharacterResponse{
			Character: character,
		}, err
	}
}
