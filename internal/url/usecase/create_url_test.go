package usecase_test

import (
	"testing"

	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/usecase"
	"github.com/matheusapostulo/url-shortener/mocks"
	"github.com/stretchr/testify/require"
)

func TestUsecaseCreateUrl(t *testing.T) {
	originalUrlInputTest := "http://www.google.com"
	shortenerUrlOutputTest := "http://localhost:8080/1"

	testCases := []struct {
		name                              string
		cacheMock                         func() *mocks.CacheRepository
		expectedCacheMockCalls            int
		urlRpMock                         func() *mocks.URLRepository
		expectedRpLongUrlMockCalls        int
		expectedRpNewAvailableIDMockCalls int
		shortenerSvMock                   func() *mocks.URLShortener
		expectedSvMockCalls               int
		input                             usecase.CreateURLInputDto
		expected                          usecase.CreateURLOutputDto
	}{
		{
			name: "Should create a shorter url",
			cacheMock: func() *mocks.CacheRepository {
				mk := mocks.NewCacheRepository(t)
				mk.On("Get", originalUrlInputTest).Return(nil, nil)
				return mk
			},
			expectedCacheMockCalls: 1,
			urlRpMock: func() *mocks.URLRepository {
				mk := mocks.NewURLRepository(t)
				mk.On("FindByLongURL", originalUrlInputTest).Return(nil, nil)
				mk.On("GetNewAvailableID").Return(1, nil)
				return mk
			},
			expectedRpLongUrlMockCalls:        1,
			expectedRpNewAvailableIDMockCalls: 1,
			shortenerSvMock: func() *mocks.URLShortener {
				mk := mocks.NewURLShortener(t)
				mk.On("ShortenURL", 1).Return(shortenerUrlOutputTest, nil)
				return mk
			},
			expectedSvMockCalls: 1,
			input: usecase.CreateURLInputDto{
				LongURL: originalUrlInputTest,
			},
			expected: usecase.CreateURLOutputDto{
				ShortURL: shortenerUrlOutputTest,
			},
		},
		{
			name: "Should return the shorter url from cache",
			cacheMock: func() *mocks.CacheRepository {
				mk := mocks.NewCacheRepository(t)
				mk.On("Get", originalUrlInputTest).Return(&domain.URL{ShortURL: shortenerUrlOutputTest}, nil)
				return mk
			},
			expectedCacheMockCalls: 1,
			urlRpMock: func() *mocks.URLRepository {
				mk := mocks.NewURLRepository(t)
				return mk
			},
			expectedRpLongUrlMockCalls:        0,
			expectedRpNewAvailableIDMockCalls: 0,
			shortenerSvMock: func() *mocks.URLShortener {
				mk := mocks.NewURLShortener(t)
				return mk
			},
			expectedSvMockCalls: 0,
			input: usecase.CreateURLInputDto{
				LongURL: originalUrlInputTest,
			},
			expected: usecase.CreateURLOutputDto{
				ShortURL: shortenerUrlOutputTest,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cacheRp := tc.cacheMock()
			urlRp := tc.urlRpMock()
			shortenerSv := tc.shortenerSvMock()

			createUrl := usecase.NewCreateURLUsecase(urlRp, cacheRp, shortenerSv)
			output, err := createUrl.Execute(tc.input)

			require.NoError(t, err)
			require.Equal(t, tc.expected.ShortURL, output.ShortURL)
			cacheRp.AssertNumberOfCalls(t, "Get", tc.expectedCacheMockCalls)
			urlRp.AssertNumberOfCalls(t, "FindByLongURL", tc.expectedRpLongUrlMockCalls)
			urlRp.AssertNumberOfCalls(t, "GetNewAvailableID", tc.expectedRpNewAvailableIDMockCalls)
			shortenerSv.AssertNumberOfCalls(t, "ShortenURL", tc.expectedSvMockCalls)
		})
	}

}
