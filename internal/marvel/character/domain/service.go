package domain

import "context"

type CharacterService interface {
	Create(ctx context.Context, character Character) (*Character, error)
}

type BasicCharacterService struct {
	repo Repository
}

func NewBasicCharacterService(repo Repository) CharacterService {
	return &BasicCharacterService{
		repo: repo,
	}
}

func (s *BasicCharacterService) Create(ctx context.Context, character Character) (*Character, error) {
	return s.repo.Save(ctx, character)
}
