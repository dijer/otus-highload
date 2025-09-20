package cache_feed

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dijer/otus-highload/backend/internal/models"
	utils_pointer "github.com/dijer/otus-highload/backend/internal/utils/pointer"
	"github.com/redis/go-redis/v9"
)

const maxFeedSize = 1000

type FeedCache struct {
	redis *redis.Client
}

func New(redis *redis.Client) *FeedCache {
	return &FeedCache{
		redis: redis,
	}
}

func makeRedisKey(userID int64) string {
	return fmt.Sprintf("feed:%d", userID)
}

func (c *FeedCache) PushToFeed(ctx context.Context, userID int64, post models.Post) error {
	key := makeRedisKey(userID)

	data, err := json.Marshal(post)
	if err != nil {
		return err
	}

	if err := c.redis.LPush(ctx, key, data).Err(); err != nil {
		return err
	}

	return c.redis.LTrim(ctx, key, 0, maxFeedSize-1).Err()
}

func (c *FeedCache) RemoveFromFeed(ctx context.Context, userID, postID int64, followers []int64) error {
	for _, followerID := range followers {
		key := makeRedisKey(followerID)

		items, err := c.redis.LRange(ctx, key, 0, -1).Result()
		if err != nil {
			continue
		}

		for _, item := range items {
			var p models.Post
			if err := json.Unmarshal([]byte(item), &p); err != nil {
				continue
			}
			if p.ID == postID {
				c.redis.LRem(ctx, key, 0, item)
			}
		}
	}

	return nil
}

func (c *FeedCache) GetFeed(ctx context.Context, userID int64, limit, offset *int64) ([]models.Post, error) {
	l := utils_pointer.ValueOrDefault(limit, 20)
	o := utils_pointer.ValueOrDefault(offset, 0)
	key := makeRedisKey(userID)

	start := int64(o)
	stop := int64(o + l - 1)

	items, err := c.redis.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	posts := make([]models.Post, 0, len(items))
	for _, item := range items {
		var p models.Post
		if err := json.Unmarshal([]byte(item), &p); err == nil {
			posts = append(posts, p)
		}
	}

	return posts, nil
}
