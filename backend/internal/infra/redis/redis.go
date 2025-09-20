package infra_redis

import (
	"context"
	"fmt"

	"github.com/dijer/otus-highload/backend/internal/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(ctx context.Context, cfg config.RedisConf) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	redisDB := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DBIndex,
	})

	if err := redisDB.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return redisDB, nil
}
