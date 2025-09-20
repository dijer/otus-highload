package service_posts

import (
	"context"

	"github.com/dijer/otus-highload/backend/internal/models"
	storage_posts "github.com/dijer/otus-highload/backend/internal/storage/posts"
)

type PostsService struct {
	storage *storage_posts.PostsStorage
}

func New(storage *storage_posts.PostsStorage) *PostsService {
	return &PostsService{
		storage: storage,
	}
}

func (s *PostsService) CreatePost(ctx context.Context, userID int64, content string) (*models.Post, error) {
	return s.storage.CreatePost(ctx, userID, content)
}

func (s *PostsService) UpdatePost(ctx context.Context, userID, postID int64, content string) error {
	return s.storage.UpdatePost(ctx, userID, postID, content)
}

func (s *PostsService) DeletePost(ctx context.Context, userID, postID int64) error {
	return s.storage.DeletePost(ctx, userID, postID)
}

func (s *PostsService) GetPost(ctx context.Context, userID, postID int64) (*models.Post, error) {
	return s.storage.GetPost(ctx, userID, postID)
}

func (s *PostsService) GetFeed(ctx context.Context, userID int64, limit, offset *int64) ([]models.Post, error) {
	return s.storage.GetFeed(ctx, userID, limit, offset)
}
