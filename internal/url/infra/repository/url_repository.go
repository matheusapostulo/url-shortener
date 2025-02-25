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
	query := "SELECT MAX(id) FROM urls"
	row := u.db.QueryRow(query)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}

func (u *URLRepositoryDatabase) FindByShortURL(shortURL string) (*domain.URL, error) {
	query := "SELECT id, long_url, short_url FROM urls WHERE short_url = ?"
	var url domain.URL

	row := u.db.QueryRow(query, shortURL)

	err := row.Scan(&url.ID, &url.LongURL, &url.ShortURL)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (u *URLRepositoryDatabase) FindByLongURL(longURL string) (*domain.URL, error) {
	query := "SELECT id, long_url, short_url FROM urls WHERE long_url = ?"
	var url domain.URL

	row := u.db.QueryRow(query, longURL)

	err := row.Scan(&url.ID, &url.LongURL, &url.ShortURL)
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (u *URLRepositoryDatabase) Save(url *domain.URL) error {
	query := "INSERT INTO urls (long_url, short_url) VALUES (?, ?)"
	result, err := u.db.Exec(query, (*url).LongURL, (*url).ShortURL)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	(*url).ID = int(id)

	return nil
}

func (u *URLRepositoryDatabase) Delete(id int) error {
	query := "DELETE FROM urls WHERE id = ?"
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
