package cache

import (
	"context"
	"fmt"
	"gin-api/config"
	"github.com/redis/go-redis/v9"
)

var client *Client

type Client struct {
	redisClient *redis.Client
	ctx         context.Context
}

func GetClient() *Client {
	return client
}

func NewClient(ctx context.Context, cfg *config.CacheConfig) error {
	if !cfg.Enable {
		return nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("could not connect to Redis: %v", err)
	}

	client = &Client{
		redisClient: rdb,
		ctx:         ctx,
	}

	return nil
}
