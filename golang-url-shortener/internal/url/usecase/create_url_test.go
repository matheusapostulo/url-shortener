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

func TestUsecaseCreateUrl(t *testing.T) {
	originalUrlInputTest := "http://www.google.com"
	shortenerUrlOutputTest := "1"

	testCases := []struct {
		name                              string
		cacheMock                         func() *mocks.CacheRepository
		expectedCacheMockCalls            int
		urlRpMock                         func() *mocks.URLRepository
		expectedRpLongUrlMockCalls        int
		expectedRpNewAvailableIDMockCalls int
		shortenerSvMock                   func() *mocks.URLShortener
		expectedSvMockCalls               int
		logPublisherMock                  func() *mocks.LogPublisherService
		expectedLogPublisherMockCalls     int
		input                             port.CreateURLInputDto
		expected                          port.CreateURLOutputDto
	}{
		{
			name: "Should create a shorter url",
			cacheMock: func() *mocks.CacheRepository {
				mk := mocks.NewCacheRepository(t)
				mk.On("Set", mock.Anything).Return(nil)
				return mk
			},
			expectedCacheMockCalls: 1,
			urlRpMock: func() *mocks.URLRepository {
				mk := mocks.NewURLRepository(t)
				mk.On("FindByLongURL", originalUrlInputTest).Return(nil, nil)
				mk.On("GetNewAvailableID").Return(1, nil)
				mk.On("Save", mock.Anything).Return(nil)
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
			logPublisherMock: func() *mocks.LogPublisherService {
				mk := mocks.NewLogPublisherService(t)
				mk.On("PublishLog", mock.Anything).Return(nil)
				return mk
			},
			expectedLogPublisherMockCalls: 1,
			input: port.CreateURLInputDto{
				LongURL: originalUrlInputTest,
			},
			expected: port.CreateURLOutputDto{
				ShortURL: shortenerUrlOutputTest,
			},
		},
		{
			name: "Should return the shorter url from database",
			cacheMock: func() *mocks.CacheRepository {
				mk := mocks.NewCacheRepository(t)
				mk.On("Set", mock.Anything).Return(nil)
				return mk
			},
			expectedCacheMockCalls: 1,
			urlRpMock: func() *mocks.URLRepository {
				mk := mocks.NewURLRepository(t)
				mk.On("FindByLongURL", originalUrlInputTest).Return(&domain.URL{ShortURL: shortenerUrlOutputTest}, nil)
				return mk
			},
			expectedRpLongUrlMockCalls:        1,
			expectedRpNewAvailableIDMockCalls: 0,
			shortenerSvMock: func() *mocks.URLShortener {
				mk := mocks.NewURLShortener(t)
				return mk
			},
			expectedSvMockCalls: 0,
			logPublisherMock: func() *mocks.LogPublisherService {
				mk := mocks.NewLogPublisherService(t)
				mk.On("PublishLog", mock.Anything).Return(nil)
				return mk
			},
			expectedLogPublisherMockCalls: 1,
			input: port.CreateURLInputDto{
				LongURL: originalUrlInputTest,
			},
			expected: port.CreateURLOutputDto{
				ShortURL: shortenerUrlOutputTest,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			urlRp := tc.urlRpMock()
			cacheRp := tc.cacheMock()
			shortenerSv := tc.shortenerSvMock()
			logPublisherSv := tc.logPublisherMock()

			createUrl := usecase.NewCreateURLUsecase(urlRp, cacheRp, shortenerSv, logPublisherSv)
			output, err := createUrl.Execute(tc.input)

			require.NoError(t, err)
			require.Equal(t, tc.expected.ShortURL, output.ShortURL)
			urlRp.AssertNumberOfCalls(t, "FindByLongURL", tc.expectedRpLongUrlMockCalls)
			urlRp.AssertNumberOfCalls(t, "GetNewAvailableID", tc.expectedRpNewAvailableIDMockCalls)
			shortenerSv.AssertNumberOfCalls(t, "ShortenURL", tc.expectedSvMockCalls)
			cacheRp.AssertNumberOfCalls(t, "Set", tc.expectedCacheMockCalls)
			logPublisherSv.AssertNumberOfCalls(t, "PublishLog", tc.expectedLogPublisherMockCalls)
		})
	}

}
