package mongodb

import (
	"context"

	"github.com/avelex/blog/internal/blog/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type blogRepository struct {
	col *mongo.Collection
}

func NewBlogRepository(database *mongo.Database) *blogRepository {
	return &blogRepository{col: database.Collection("blog")}
}

func (r *blogRepository) GetPost(ctx context.Context, id string) (entity.Post, error) {
	filter := bson.D{{"id", id}}

	var result entity.Post
	err := r.col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return entity.Post{}, err
	}

	return result, nil
}
func (r *blogRepository) GetPosts(ctx context.Context) ([]entity.Post, error) {
	cur, err := r.col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	result := make([]entity.Post, 0)

	for cur.Next(ctx) {
		var post entity.Post
		if err := cur.Decode(&post); err != nil {
			return nil, err
		}

		result = append(result, post)
	}

	if err := cur.Close(ctx); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *blogRepository) CreatePost(ctx context.Context, post entity.Post) error {
	_, err := r.col.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	return nil
}
