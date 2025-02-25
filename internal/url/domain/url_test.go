package domain_test

import (
	"testing"

	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/stretchr/testify/require"
)

func TestNewURL(t *testing.T) {
	t.Run("Should create a new URL", func(t *testing.T) {
		url, err := domain.NewURL(1, "http://www.google.com", "http://localhost:8080/1")

		require.NoError(t, err)
		require.Equal(t, 1, url.ID)
		require.Equal(t, "http://www.google.com", url.LongURL)
		require.Equal(t, "http://localhost:8080/1", url.ShortURL)
	})
}
