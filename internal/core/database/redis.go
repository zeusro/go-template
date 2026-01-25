package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"zeusro.com/hermes/internal/core/config"
	"zeusro.com/hermes/internal/core/logprovider"
)

// RedisClient wraps Redis client
type RedisClient struct {
	*redis.Client
}

// NewRedisClient creates a new Redis client connection
func NewRedisClient(cfg config.Config, log logprovider.Logger) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
		PoolSize: cfg.Redis.PoolSize,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Infof("Successfully connected to Redis at %s:%d", cfg.Redis.Host, cfg.Redis.Port)

	return &RedisClient{Client: rdb}, nil
}

// Close closes the Redis connection
func (r *RedisClient) Close() error {
	return r.Client.Close()
}
