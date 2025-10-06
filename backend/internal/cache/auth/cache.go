package cache_auth

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthCache struct {
	cache *redis.Client
}

func New(cache *redis.Client) *AuthCache {
	return &AuthCache{
		cache: cache,
	}
}

func (c *AuthCache) SaveSession(ctx context.Context, uuid string, userID int64, ttl time.Duration) error {
	key := c.generateKey(uuid)
	return c.cache.SetEx(ctx, key, userID, ttl).Err()
}

func (c *AuthCache) GetSession(ctx context.Context, uuid string) (int64, error) {
	key := c.generateKey(uuid)
	val, err := c.cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (c *AuthCache) DeleteSession(ctx context.Context, uuid string) error {
	key := c.generateKey(uuid)
	return c.cache.Del(ctx, key).Err()
}

func (c *AuthCache) generateKey(uuid string) string {
	return "sess:" + uuid
}
