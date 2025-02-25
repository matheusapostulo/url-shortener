package repository

import "github.com/matheusapostulo/url-shortener/internal/url/domain"

func NewCacheRepositoryRedis() *CacheRepositoryRedis {
	return &CacheRepositoryRedis{}
}

type CacheRepositoryRedis struct {
}

func (c *CacheRepositoryRedis) Get(key string) (*domain.URL, error) {
	return nil, nil
}

func (c *CacheRepositoryRedis) Set(key string, url *domain.URL) error {
	return nil
}
