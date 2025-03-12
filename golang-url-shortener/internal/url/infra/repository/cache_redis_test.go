package repository

import (
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/stretchr/testify/require"
)

func TestCacheRedisGet(t *testing.T) {
	t.Run("Should return a url", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()
		mock.ExpectGet("short").SetVal(`{"id":1,"long_url":"long","short_url":"short"}`)

		cache := NewCacheRepositoryRedis(rdb)

		url, err := cache.Get("short")

		expectedURL := &domain.URL{
			ID:       1,
			LongURL:  "long",
			ShortURL: "short",
		}

		mock.ExpectationsWereMet()
		require.NoError(t, err)
		require.Equal(t, expectedURL, url)
	})
}

func TestCacheRedisSet(t *testing.T) {
	t.Run("Should set a url", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()

		mock.ExpectSet("short", `{"id":1,"long_url":"long","short_url":"short"}`, 1*time.Hour).SetVal("OK")

		url := &domain.URL{
			ID:       1,
			LongURL:  "long",
			ShortURL: "short",
		}

		cache := NewCacheRepositoryRedis(rdb)

		err := cache.Set(url)

		require.NoError(t, err)
	})

	t.Run("Should return an error", func(t *testing.T) {
		rdb, mock := redismock.NewClientMock()
		url := &domain.URL{
			ID:       1,
			LongURL:  "long",
			ShortURL: "short",
		}

		mock.ExpectSet("short", url, 1*time.Hour).SetErr(nil)

		cache := NewCacheRepositoryRedis(rdb)

		err := cache.Set(url)

		require.Error(t, err)
	})
}
