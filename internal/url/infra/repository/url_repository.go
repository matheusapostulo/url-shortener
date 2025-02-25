package repository

import (
	"database/sql"

	"github.com/matheusapostulo/url-shortener/internal/url/domain"
)

func NewURLRepositoryDatabase(db *sql.DB) *URLRepositoryDatabase {
	return &URLRepositoryDatabase{
		db: db,
	}
}

type URLRepositoryDatabase struct {
	db *sql.DB
}

func (u *URLRepositoryDatabase) GetNewAvailableID() (int, error) {
	return 0, nil
}

func (u *URLRepositoryDatabase) FindByShortURL(shortURL string) (*domain.URL, error) {
	return &domain.URL{}, nil
}

func (u *URLRepositoryDatabase) FindByLongURL(longURL string) (*domain.URL, error) {
	return &domain.URL{}, nil
}

func (u *URLRepositoryDatabase) Save(url *domain.URL) error {
	return nil
}

func (u *URLRepositoryDatabase) Delete(id int) error {
	return nil
}
