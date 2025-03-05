package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	httpPkg "github.com/matheusapostulo/url-shortener/internal/url/infra/http"
	"github.com/matheusapostulo/url-shortener/internal/url/infra/repository"
	"github.com/matheusapostulo/url-shortener/internal/url/infra/service"
	"github.com/matheusapostulo/url-shortener/internal/url/usecase"
	"github.com/redis/go-redis/v9"
)

func main() {
	mysqlCfg := mysql.Config{
		User:      "root",
		Passwd:    "root",
		Net:       "tcp",
		Addr:      "localhost:3307",
		DBName:    "urls",
		ParseTime: true,
	}

	// db
	db, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// cache
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	// Repositories
	urlRp := repository.NewURLRepositoryDatabase(db)
	cacheRp := repository.NewCacheRepositoryRedis(rdb)

	// Services
	shortenerSv := service.NewURLShortenerBase62()

	// Usecases
	createURLUsecase := usecase.NewCreateURLUsecase(urlRp, cacheRp, shortenerSv)
	redirectURLUsecase := usecase.NewRedirectURLUsecase(cacheRp, urlRp)

	// Handlers
	urlHandler := httpPkg.NewURLHandler(createURLUsecase, redirectURLUsecase)

	rt := chi.NewRouter()
	rt.Use(middleware.Logger)

	rt.Route("/api/v1", func(rt chi.Router) {
		rt.Post("/shorten", urlHandler.CreateURL)
		rt.Get("/{short-url}", urlHandler.RedirectURL)
	})

	fmt.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", rt)
	if err != nil {
		panic(err)
	}
}
