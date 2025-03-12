package repository_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/infra/repository"
	"github.com/stretchr/testify/require"
)

var (
	longURL  = "http://www.google.com"
	shortURL = "http://localhost:8080/1"
)

func TestURLRepositoryGetNewAvailableID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db2, _, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	query := "SELECT MAX(id) FROM url"

	t.Run("Should return the next available ID", func(t *testing.T) {
		mock.ExpectQuery(query).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		repo := repository.NewURLRepositoryDatabase(db, db2)

		id, err := repo.GetNewAvailableID()

		require.NoError(t, err)
		require.Equal(t, 2, id)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("case 2: error - Error retrieving the next available ID", func(t *testing.T) {
		mock.ExpectQuery(query).
			WillReturnError(sql.ErrConnDone)

		rp := repository.NewURLRepositoryDatabase(db, db2)
		_, err := rp.GetNewAvailableID()

		require.Error(t, err)
	})
}

func TestURLRepositoryFindByShortURL(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db2, _, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	require.NoError(t, err)
	defer db.Close()

	url, _ := domain.NewURL(1, longURL, shortURL)
	query := "SELECT id, long_url, short_url FROM url WHERE short_url = ?"

	t.Run("Should return a URL by short URL", func(t *testing.T) {
		mock.ExpectQuery(query).
			WithArgs(url.ShortURL).
			WillReturnRows(sqlmock.NewRows([]string{"id", "long_url", "short_url"}).
				AddRow(url.ID, url.LongURL, url.ShortURL))

		repo := repository.NewURLRepositoryDatabase(db, db2)

		result, err := repo.FindByShortURL(url.ShortURL)

		require.NoError(t, err)
		require.Equal(t, url, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("case 2: error - Error retrieving the URL by short URL", func(t *testing.T) {
		mock.ExpectQuery(query).
			WithArgs(url.ShortURL).
			WillReturnError(sql.ErrConnDone)

		rp := repository.NewURLRepositoryDatabase(db, db2)
		_, err := rp.FindByShortURL(url.ShortURL)

		require.Error(t, err)
	})
}

func TestURLRepositoryFindByLongURL(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db2, _, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	url, _ := domain.NewURL(1, longURL, shortURL)
	query := "SELECT id, long_url, short_url FROM url WHERE long_url = ?"

	t.Run("Should return a URL by long URL", func(t *testing.T) {
		mock.ExpectQuery(query).
			WithArgs(url.LongURL).
			WillReturnRows(sqlmock.NewRows([]string{"id", "long_url", "short_url"}).
				AddRow(url.ID, url.LongURL, url.ShortURL))

		repo := repository.NewURLRepositoryDatabase(db, db2)

		result, err := repo.FindByLongURL(url.LongURL)

		require.NoError(t, err)
		require.Equal(t, url, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("case 2: error - Error retrieving the URL by long URL", func(t *testing.T) {
		mock.ExpectQuery(query).
			WithArgs(url.LongURL).
			WillReturnError(sql.ErrConnDone)

		rp := repository.NewURLRepositoryDatabase(db, db2)
		_, err := rp.FindByLongURL(url.LongURL)

		require.Error(t, err)
	})
}

func TestURLRepositorySave(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db2, _, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	require.NoError(t, err)
	defer db.Close()

	url, _ := domain.NewURL(1, longURL, shortURL)
	query := "INSERT INTO url (long_url, short_url) VALUES (?, ?)"

	t.Run("Should save a URL", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(url.LongURL, url.ShortURL).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := repository.NewURLRepositoryDatabase(db, db2)

		err := repo.Save(url)

		require.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("case 2: error - Error retrieving the last inserted ID", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(url.LongURL, url.ShortURL).
			WillReturnError(sql.ErrConnDone)

		rp := repository.NewURLRepositoryDatabase(db, db2)
		err := rp.Save(url)

		require.Error(t, err)
	})

	t.Run("case 3: error - Error retrieving the last inserted ID", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(url.LongURL, url.ShortURL).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))

		rp := repository.NewURLRepositoryDatabase(db, db2)
		err := rp.Save(url)

		require.Error(t, err)
	})
}

func TestURLRepositoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db2, _, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	url, _ := domain.NewURL(1, longURL, shortURL)
	query := "DELETE FROM url WHERE id = ?"

	t.Run("Should delete a URL", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(url.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := repository.NewURLRepositoryDatabase(db, db2)

		err := repo.Delete(url.ID)

		require.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("case 2: error - Error deleting the URL", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(url.ID).
			WillReturnError(sql.ErrConnDone)

		rp := repository.NewURLRepositoryDatabase(db, db2)
		err := rp.Delete(url.ID)

		require.Error(t, err)
	})
}
