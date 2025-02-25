package repository

import "github.com/matheusapostulo/url-shortener/internal/url/domain"

func NewCacheRepositoryRedis() *CacheRepositoryRedis {
	return &CacheRepositoryRedis{}
}

type CacheRepositoryRedis struct {
}

func (c *CacheRepositoryRedis) Get(key string) (*domain.URL, error) {
	return &domain.URL{}, nil
}

func (c *CacheRepositoryRedis) Set(key string, url string) error {
	return nil
}
