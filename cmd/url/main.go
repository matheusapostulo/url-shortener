package main

import (
	"database/sql"

	"github.com/matheusapostulo/url-shortener/internal/url/infra/repository"
	"github.com/matheusapostulo/url-shortener/internal/url/infra/service"
	"github.com/matheusapostulo/url-shortener/internal/url/usecase"
)

func main() {
	db, err := sql.Open("mysql	", "test_user:test_password@tcp(mysql:3307)/test_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Repositories
	urlRp := repository.NewURLRepositoryDatabase(db)
	cacheRp := repository.NewCacheRepositoryRedis()

	// Services
	shortenerSv := service.NewURLShortenerBase62()

	// Usecases
	urlUsecase := usecase.NewCreateURLUsecase(urlRp, cacheRp, shortenerSv)

}
