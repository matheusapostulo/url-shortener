package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/redis/go-redis/v9"
)

func NewCacheRepositoryRedis(client *redis.Client) *CacheRepositoryRedis {
	return &CacheRepositoryRedis{
		client: client,
		ctx:    context.Background(),
	}
}

type CacheRepositoryRedis struct {
	client *redis.Client
	ctx    context.Context
}

func (c *CacheRepositoryRedis) Get(key string) (*domain.URL, error) {
	result, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, domain.ErrURLNotFound
		}
		return nil, err
	}

	var url domain.URL
	err = json.Unmarshal([]byte(result), &url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &url, nil
}

func (c *CacheRepositoryRedis) Set(url *domain.URL) error {
	data, err := json.Marshal(url)
	if err != nil {
		return err
	}

	err = c.client.Set(c.ctx, url.ShortURL, string(data), 1*time.Hour).Err()
	if err != nil {
		fmt.Println("Err set", err)
		return err
	}

	return nil
}
