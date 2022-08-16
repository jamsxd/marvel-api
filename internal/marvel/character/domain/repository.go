package domain

import "context"

type Repository interface {
	Save(ctx context.Context, character Character) (*Character, error)
}
