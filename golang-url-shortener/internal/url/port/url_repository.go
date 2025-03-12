package port

import (
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
)

type URLRepository interface {
	GetNewAvailableID() (int, error)
	FindByShortURL(shortURL string) (*domain.URL, error)
	FindByLongURL(longURL string) (*domain.URL, error)
	Save(url *domain.URL) error
	Delete(id int) error
}
