package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/avelex/blog/internal/blog/entity"
	"go.uber.org/zap"
)

type BlogRepository interface {
	GetPost(ctx context.Context, id string) (entity.Post, error)
	GetPosts(ctx context.Context) ([]entity.Post, error)
	CreatePost(ctx context.Context, post entity.Post) error
}

type blogUsecase struct {
	logger *zap.SugaredLogger
	repo   BlogRepository
}

func NewUsecase(logger *zap.SugaredLogger, repo BlogRepository) *blogUsecase {
	return &blogUsecase{
		logger: logger,
		repo:   repo,
	}
}

func (u *blogUsecase) GetPosts(ctx context.Context) ([]entity.Post, error) {
	ctxGet, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	posts, err := u.repo.GetPosts(ctxGet)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (u *blogUsecase) GetPost(ctx context.Context, id string) (entity.Post, error) {
	ctxGet, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	return u.repo.GetPost(ctxGet, id)
}

func (u *blogUsecase) CreatePost(ctx context.Context) error {
	ctxGet, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	for i := 0; i < 10; i++ {
		id := strconv.Itoa(i)
		if err := u.repo.CreatePost(ctxGet, entity.Post{
			ID:          id,
			Title:       "Title " + id,
			Description: "Desc " + id,
			Body:        "Body",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}); err != nil {
			return err
		}
	}

	return nil
}
