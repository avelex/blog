package mongodb

import (
	"context"

	"github.com/avelex/blog/internal/blog/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type blogRepository struct {
	col *mongo.Collection
}

func NewBlogRepository(database *mongo.Database) *blogRepository {
	return &blogRepository{col: database.Collection("blog")}
}

func (r *blogRepository) GetPost(ctx context.Context, id string) (entity.Post, error) {
	return entity.Post{}, nil
}
func (r *blogRepository) GetPosts(ctx context.Context) ([]entity.Post, error) {
	return nil, nil
}
