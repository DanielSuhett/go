package post

import (
	"context"
	"postgres/entities"
)

type PostRepository interface {
	Create(ctx context.Context, post *entities.Post) error
	Update(ctx context.Context, id int, post *entities.Post) error
	Get(ctx context.Context, id int) (*entities.Post, error)
	GetAll(ctx context.Context) ([]entities.Post, error)
}
