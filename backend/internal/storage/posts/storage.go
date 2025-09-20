package storage_posts

import (
	"context"
	"database/sql"

	cache_feed "github.com/dijer/otus-highload/backend/internal/cache/feed"
	infra_database "github.com/dijer/otus-highload/backend/internal/infra/database"
	"github.com/dijer/otus-highload/backend/internal/logger"
	"github.com/dijer/otus-highload/backend/internal/models"
)

type PostsStorage struct {
	dbRouter infra_database.DBRouter
	cache    *cache_feed.FeedCache
	log      logger.Logger
}

func New(dbRouter infra_database.DBRouter, cache *cache_feed.FeedCache, log logger.Logger) *PostsStorage {
	return &PostsStorage{
		dbRouter: dbRouter,
		cache:    cache,
		log:      log,
	}
}

func (s *PostsStorage) CreatePost(ctx context.Context, userID int64, content string) (*models.Post, error) {
	var post models.Post
	err := s.dbRouter.QueryRow(ctx,
		`INSERT INTO posts (user_id, content, created_at, updated_at)
				VALUES ($1, $2, NOW(), NOW())
				RETURNING id, userId, content, createdAt, updatedAt`,
		userID, content,
	).Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if err := s.cache.PushToFeed(ctx, userID, post); err != nil {
		s.log.Warn("failed push to feed")
	}

	return &post, nil
}

func (s *PostsStorage) UpdatePost(ctx context.Context, userID, postID int64, content string) error {
	_, err := s.dbRouter.Exec(ctx,
		`UPDATE posts SET content = $1, updated_at = NOW()
				WHERE id = $2 AND user_id = $3`,
		content, postID, userID,
	)

	return err
}

func (s *PostsStorage) DeletePost(ctx context.Context, userID, postID int64) error {
	_, err := s.dbRouter.Exec(ctx,
		`DELETE FROM posts WHERE id = $1 AND user_id = $2`,
		postID, userID,
	)
	if err != nil {
		return err
	}

	followers, err := s.GetFollowers(ctx, userID)
	if err != nil {
		return err
	}

	if err := s.cache.RemoveFromFeed(ctx, userID, postID, followers); err != nil {
		s.log.Warn("failed push to feed")
	}

	return nil
}

func (s *PostsStorage) GetPost(ctx context.Context, userID, postID int64) (*models.Post, error) {
	var p models.Post
	row := s.dbRouter.QueryRow(ctx,
		`SELECT id, user_id, content, created_at, updated_at
				FROM posts WHERE id = $1`,
		postID,
	)

	if err := row.Scan(&p.ID, &p.UserID, &p.Content, &p.CreatedAt, &p.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &p, nil
}

func (s *PostsStorage) GetFeed(ctx context.Context, userID int64, limit, offset *int64) ([]models.Post, error) {
	return s.cache.GetFeed(ctx, userID, limit, offset)
}

func (s *PostsStorage) GetFollowers(ctx context.Context, userID int64) ([]int64, error) {
	rows, err := s.dbRouter.Query(ctx,
		`SELECT user_id FROM follows WHERE friend_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []int64
	for rows.Next() {
		var follower int64
		if err := rows.Scan(&follower); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}
