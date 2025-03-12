package usecase_test

import (
	"testing"

	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/port"
	"github.com/matheusapostulo/url-shortener/internal/url/usecase"
	"github.com/matheusapostulo/url-shortener/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRedirectURL(t *testing.T) {
	shortUrlInputTest := "1"
	longUrlOutputTest := "http://www.google.com"

	testCases := []struct {
		name                       string
		cacheMock                  func() *mocks.CacheRepository
		expectedCacheGetMockCalls  int
		expectedCacheSetMockCalls  int
		urlRpMock                  func() *mocks.URLRepository
		expectedRpLongUrlMockCalls int
		input                      port.RedirectURLInputDto
		expected                   port.RedirectURLOutputDto
		expectedErr                error
	}{
		{
			name: "Should return a long URL from cache",
			cacheMock: func() *mocks.CacheRepository {
				mk := mocks.NewCacheRepository(t)
				mk.On("Get", mock.Anything).Return(&domain.URL{
					LongURL: longUrlOutputTest,
				}, nil)
				return mk
			},
			expectedCacheGetMockCalls: 1,
			expectedCacheSetMockCalls: 0,
			urlRpMock: func() *mocks.URLRepository {
				return mocks.NewURLRepository(t)
			},
			expectedRpLongUrlMockCalls: 0,
			input: port.RedirectURLInputDto{
				ShortURL: shortUrlInputTest,
			},
			expected: port.RedirectURLOutputDto{
				LongURL: longUrlOutputTest,
			},
			expectedErr: nil,
		},
		{
			name: "Should return a long URL from database",
			cacheMock: func() *mocks.CacheRepository {
				mk := mocks.NewCacheRepository(t)
				mk.On("Get", mock.Anything).Return(nil, nil)
				mk.On("Set", mock.Anything).Return(nil)
				return mk
			},
			expectedCacheGetMockCalls: 1,
			expectedCacheSetMockCalls: 1,
			urlRpMock: func() *mocks.URLRepository {
				mk := mocks.NewURLRepository(t)
				mk.On("FindByShortURL", shortUrlInputTest).Return(&domain.URL{
					LongURL: longUrlOutputTest,
				}, nil)
				return mk
			},
			expectedRpLongUrlMockCalls: 1,
			input: port.RedirectURLInputDto{
				ShortURL: shortUrlInputTest,
			},
			expected: port.RedirectURLOutputDto{
				LongURL: longUrlOutputTest,
			},
			expectedErr: nil,
		},
		{
			name: "Should return an error when the short URL is not found",
			cacheMock: func() *mocks.CacheRepository {
				mk := mocks.NewCacheRepository(t)
				mk.On("Get", mock.Anything).Return(nil, nil)
				return mk
			},
			expectedCacheGetMockCalls: 1,
			expectedCacheSetMockCalls: 0,
			urlRpMock: func() *mocks.URLRepository {
				mk := mocks.NewURLRepository(t)
				mk.On("FindByShortURL", shortUrlInputTest).Return(nil, nil)
				return mk
			},
			expectedRpLongUrlMockCalls: 1,
			input: port.RedirectURLInputDto{
				ShortURL: shortUrlInputTest,
			},
			expected:    port.RedirectURLOutputDto{},
			expectedErr: domain.ErrURLNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cacheRepository := tc.cacheMock()
			urlRepository := tc.urlRpMock()

			redirectURL := usecase.NewRedirectURLUsecase(cacheRepository, urlRepository)
			output, err := redirectURL.Execute(tc.input)

			require.ErrorIs(t, err, tc.expectedErr)
			require.Equal(t, tc.expected, output)

			cacheRepository.AssertExpectations(t)
			urlRepository.AssertExpectations(t)
		})
	}
}
