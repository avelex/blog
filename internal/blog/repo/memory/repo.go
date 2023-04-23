package memory

import (
	"context"
	"errors"
	"sort"
	"strings"
	"sync"

	"github.com/avelex/blog/internal/blog/entity"
)

type memoryRepository struct {
	mutex sync.RWMutex
	posts map[string]entity.Post
}

func NewRepository() *memoryRepository {
	return &memoryRepository{
		mutex: sync.RWMutex{},
		posts: make(map[string]entity.Post),
	}
}

func (r *memoryRepository) GetPost(ctx context.Context, id string) (entity.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	post, ok := r.posts[id]
	if !ok {
		return entity.Post{}, errors.New("not exists")
	}

	return post, nil
}

func (r *memoryRepository) GetPosts(ctx context.Context) ([]entity.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := make([]entity.Post, 0, len(r.posts))
	for _, p := range r.posts {
		result = append(result, p)
	}

	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(result[i].ID, result[j].ID) == -1
	})

	return result, nil
}

func (r *memoryRepository) CreatePost(ctx context.Context, post entity.Post) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.posts[post.ID] = post

	return nil
}
