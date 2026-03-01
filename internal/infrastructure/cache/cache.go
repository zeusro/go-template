package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeusro/go-template/internal/core/database"
)

// Cache provides caching functionality using Redis
type Cache struct {
	redis *database.RedisClient
}

// NewCache creates a new cache instance
func NewCache(redis *database.RedisClient) *Cache {
	return &Cache{redis: redis}
}

// Get retrieves a value from cache
func (c *Cache) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(val), dest); err != nil {
		return fmt.Errorf("failed to unmarshal cached value: %w", err)
	}

	return nil
}

// Set stores a value in cache with expiration
func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	return c.redis.Set(ctx, key, data, expiration).Err()
}

// Delete removes a key from cache
func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.redis.Del(ctx, key).Err()
}

// Exists checks if a key exists in cache
func (c *Cache) Exists(ctx context.Context, key string) (bool, error) {
	count, err := c.redis.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetOrSet retrieves a value from cache, or sets it if not found
func (c *Cache) GetOrSet(ctx context.Context, key string, dest interface{}, expiration time.Duration, setter func() (interface{}, error)) error {
	// Try to get from cache
	err := c.Get(ctx, key, dest)
	if err == nil {
		return nil
	}

	// If not found, call setter
	value, err := setter()
	if err != nil {
		return err
	}

	// Set in cache
	if err := c.Set(ctx, key, value, expiration); err != nil {
		return err
	}

	// Unmarshal to dest
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}
