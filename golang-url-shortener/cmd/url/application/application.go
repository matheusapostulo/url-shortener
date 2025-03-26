package application

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/matheusapostulo/url-shortener/cmd/url/dependencies"
	httpPkg "github.com/matheusapostulo/url-shortener/internal/url/infra/http"
	"github.com/matheusapostulo/url-shortener/internal/url/infra/repository"
	"github.com/matheusapostulo/url-shortener/internal/url/infra/service"
	"github.com/matheusapostulo/url-shortener/internal/url/usecase"
)

func Run(depend dependencies.Dependencies) {
	// Repositories
	urlRp := repository.NewURLRepositoryDatabase(depend.DB)
	cacheRp := repository.NewCacheRepositoryRedis(depend.CacheClient)

	// Services
	shortenerSv := service.NewURLShortenerBase62()
	logPublisherSv := service.NewLogPublisherService(depend.RabbitMQConn)

	// Usecases
	createURLUsecase := usecase.NewCreateURLUsecase(urlRp, cacheRp, shortenerSv, logPublisherSv)
	redirectURLUsecase := usecase.NewRedirectURLUsecase(cacheRp, urlRp)

	// Handlers
	urlHandler := httpPkg.NewURLHandler(createURLUsecase, redirectURLUsecase)

	rt := chi.NewRouter()

	rt.Use(middleware.Logger)
	rt.Route("/api/v1", func(rt chi.Router) {
		rt.Post("/shorten", urlHandler.CreateURL)
		rt.Get("/{short-url}", urlHandler.RedirectURL)
	})

	fmt.Println("Server running on port 3000")
	err := http.ListenAndServe(":3000", rt)
	if err != nil {
		panic(err)
	}
}
