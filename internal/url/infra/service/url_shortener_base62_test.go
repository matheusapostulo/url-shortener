package service_test

import (
	"testing"

	"github.com/matheusapostulo/url-shortener/internal/url/infra/service"
	"github.com/stretchr/testify/require"
)

func TestShorterUrl(t *testing.T) {
	t.Run("Should shorten a url whit a base62 algorithm", func(t *testing.T) {
		input := 11157
		shortener := service.NewURLShortenerBase62()

		output, _ := shortener.ShortenURL(input)

		require.Equal(t, "2TX", output)
	})
}
