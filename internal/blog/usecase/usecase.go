package usecase

import (
	"context"

	"github.com/avelex/blog/internal/blog/entity"
)

type BlogRepository interface {
	GetPost(ctx context.Context, id string) (entity.Post, error)
	GetPosts(ctx context.Context) ([]entity.Post, error)
}

type blogUsecase struct {
	repo BlogRepository
}

func NewUsecase(repo BlogRepository) *blogUsecase {
	return &blogUsecase{
		repo: repo,
	}
}

func (u *blogUsecase) GetPosts(ctx context.Context) ([]entity.Post, error) {
	return nil, nil
}

func (u *blogUsecase) GetPost(ctx context.Context, id string) (entity.Post, error) {
	return entity.Post{}, nil
}
