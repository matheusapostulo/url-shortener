package usecase_test

import (
	"testing"

	"github.com/matheusapostulo/url-shortener/internal/url/usecase"
	"github.com/matheusapostulo/url-shortener/mocks"
	"github.com/stretchr/testify/require"
)

type CreateUrlInputDto struct {
	LongUrl string
}

func TestUsecaseCreateUrl(t *testing.T) {
	t.Run("Should create a shorter url", func(t *testing.T) {
		input := usecase.CreateURLInputDto{
			LongURL: "http://www.google.com",
		}

		urlRp := mocks.NewURLRepository(t)
		cacheRp := mocks.NewCacheRepository(t)
		shortenerSv := mocks.NewURLShortener(t)
		createUrl := usecase.NewCreateURLUsecase(urlRp, cacheRp, shortenerSv)
		output, err := createUrl.Execute(input)

		expectedOutput := usecase.CreateURLOutputDto{
			ShortURL: "http://localhost:8080/1",
		}
		require.NoError(t, err)
		require.Equal(t, expectedOutput.ShortURL, output.ShortURL)
	})
}
