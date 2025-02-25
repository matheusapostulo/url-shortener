package port

import "github.com/matheusapostulo/url-shortener/internal/url/domain"

type CacheRepository interface {
	Get(key string) (*domain.URL, error)
	Set(key string, url *domain.URL) error
}
