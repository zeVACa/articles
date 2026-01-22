package service

import (
	"articles/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	client *redis.Client
	ttl    time.Duration
}

func NewCacheService(client *redis.Client, ttl time.Duration) *CacheService {
	return &CacheService{
		client: client,
		ttl:    ttl,
	}
}

func (c *CacheService) GetPost(ctx context.Context, id int64) (*models.Post, error) {
	key := fmt.Sprintf("post:%d", id)

	data, err := c.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("redis get error: %w", err)
	}

	var post models.Post
	if err := json.Unmarshal(data, &post); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &post, nil
}

func (c *CacheService) SetPost(ctx context.Context, post models.Post) error {
	key := fmt.Sprintf("post:%d", post.ID)

	data, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	if err = c.client.Set(ctx, key, data, c.ttl).Err(); err != nil {
		return fmt.Errorf("redis set error: %w", err)
	}

	return nil
}

func (c *CacheService) DeletePost(ctx context.Context, id int64) error {
	key := fmt.Sprintf("post:%d", id)
	return c.client.Del(ctx, key).Err()
}

func (c *CacheService) GetAllPosts(ctx context.Context) ([]models.Post, error) {
	key := "posts:all"

	data, err := c.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("redis get error: %w", err)
	}

	var posts []models.Post
	if err := json.Unmarshal(data, &posts); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return posts, nil
}

func (c *CacheService) SetAllPosts(ctx context.Context, posts []models.Post) error {
	key := "posts:all"

	data, err := json.Marshal(posts)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	if err = c.client.Set(ctx, key, data, c.ttl).Err(); err != nil {
		return fmt.Errorf("redis set error: %w", err)
	}

	return nil
}

func (c *CacheService) InvalidateAllPosts(ctx context.Context) error {
	key := "posts:all"
	return c.client.Del(ctx, key).Err()
}
