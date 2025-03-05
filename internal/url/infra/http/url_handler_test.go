package http_test

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/infra/http"
	"github.com/matheusapostulo/url-shortener/internal/url/port"
	"github.com/matheusapostulo/url-shortener/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUrlHandlerCreateURL(t *testing.T) {
	longUrl := `
		{
			"long_url": "http://www.google.com"
		}
	`

	testCases := []struct {
		name                   string
		createURLUsecaseMock   func() *mocks.CreateURLUsecase
		redirectURLUsecaseMock func() *mocks.RedirectURLUsecase
		input                  string
		expectedBody           string
		expectedStatus         int
		expectedError          error
	}{
		{
			name: "CreateURL with valid URL",
			createURLUsecaseMock: func() *mocks.CreateURLUsecase {
				mk := mocks.NewCreateURLUsecase(t)
				mk.On("Execute", mock.Anything).Return(port.CreateURLOutputDto{
					ShortURL: "1",
				}, nil)
				return mk
			},
			redirectURLUsecaseMock: func() *mocks.RedirectURLUsecase {
				return mocks.NewRedirectURLUsecase(t)
			},
			input:          longUrl,
			expectedBody:   `{"short_url":"1"}`,
			expectedStatus: 201,
			expectedError:  nil,
		},
		{
			name: "CreateURL usecase throws error",
			createURLUsecaseMock: func() *mocks.CreateURLUsecase {
				mk := mocks.NewCreateURLUsecase(t)
				mk.On("Execute", mock.Anything).Return(port.CreateURLOutputDto{}, domain.ErrInternalServerError)
				return mk
			},
			redirectURLUsecaseMock: func() *mocks.RedirectURLUsecase {
				return mocks.NewRedirectURLUsecase(t)
			},
			input: longUrl,
			expectedBody: `{
				"status": "Internal Server Error",
				"message": "internal server error"
			}`,
			expectedStatus: 500,
			expectedError:  domain.ErrInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createURLUsecaseMock := tc.createURLUsecaseMock()
			redirectURLUsecaseMock := tc.redirectURLUsecaseMock()

			urlHandler := http.NewURLHandler(createURLUsecaseMock, redirectURLUsecaseMock)

			request := httptest.NewRequest("POST", "/shorten", strings.NewReader(tc.input))
			response := httptest.NewRecorder()

			urlHandler.CreateURL(response, request)

			require.Equal(t, tc.expectedStatus, response.Code)
			require.JSONEq(t, tc.expectedBody, response.Body.String())
			require.ErrorIs(t, tc.expectedError, tc.expectedError)
		})
	}
}

func TestUrlHandlerRedirectURL(t *testing.T) {
	testCases := []struct {
		name                   string
		createURLUsecaseMock   func() *mocks.CreateURLUsecase
		redirectURLUsecaseMock func() *mocks.RedirectURLUsecase
		input                  string
		expectedBody           string
		expectedStatus         int
		expectedError          error
	}{
		{
			name: "RedirectURL with valid URL",
			createURLUsecaseMock: func() *mocks.CreateURLUsecase {
				return mocks.NewCreateURLUsecase(t)
			},
			redirectURLUsecaseMock: func() *mocks.RedirectURLUsecase {
				mk := mocks.NewRedirectURLUsecase(t)
				mk.On("Execute", mock.Anything).Return(port.RedirectURLOutputDto{
					LongURL: "http://www.google.com",
				}, nil)
				return mk
			},
			input:          "1",
			expectedStatus: 301,
			expectedError:  nil,
		},
		{
			name: "RedirectURL usecase throws not found error",
			createURLUsecaseMock: func() *mocks.CreateURLUsecase {
				return mocks.NewCreateURLUsecase(t)
			},
			redirectURLUsecaseMock: func() *mocks.RedirectURLUsecase {
				mk := mocks.NewRedirectURLUsecase(t)
				mk.On("Execute", mock.Anything).Return(port.RedirectURLOutputDto{}, domain.ErrURLNotFound)
				return mk
			},
			input:          "1",
			expectedBody:   `{"status":"Not Found","message":"URL not found"}`,
			expectedStatus: 404,
			expectedError:  domain.ErrURLNotFound,
		},
		{
			name: "RedirectURL usecase throws internal server error",
			createURLUsecaseMock: func() *mocks.CreateURLUsecase {
				return mocks.NewCreateURLUsecase(t)
			},
			redirectURLUsecaseMock: func() *mocks.RedirectURLUsecase {
				mk := mocks.NewRedirectURLUsecase(t)
				mk.On("Execute", mock.Anything).Return(port.RedirectURLOutputDto{}, domain.ErrInternalServerError)
				return mk
			},
			input:          "1",
			expectedBody:   `{"status":"Internal Server Error","message":"internal server error"}`,
			expectedStatus: 500,
			expectedError:  domain.ErrInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createURLUsecaseMock := tc.createURLUsecaseMock()
			redirectURLUsecaseMock := tc.redirectURLUsecaseMock()

			urlHandler := http.NewURLHandler(createURLUsecaseMock, redirectURLUsecaseMock)

			request := httptest.NewRequest("GET", "/1", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("short-url", "1")
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()

			urlHandler.RedirectURL(response, request)

			if tc.expectedError != nil {
				require.JSONEq(t, tc.expectedBody, response.Body.String())
			}
			require.Equal(t, tc.expectedStatus, response.Code)
			require.ErrorIs(t, tc.expectedError, tc.expectedError)
		})
	}
}
