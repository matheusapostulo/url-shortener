package dependencies

import (
	"database/sql"

	"github.com/matheusapostulo/url-shortener/cmd/url/config"
	"github.com/redis/go-redis/v9"
)

type Dependencies struct {
	WriteDB     *sql.DB
	ReadDB      *sql.DB
	CacheClient *redis.Client
}

func BuildDependencies() (*Dependencies, error) {
	//db
	writeDb, readDb, err := config.BuildDb()
	if err != nil {
		return nil, err
	}

	// cache
	client := config.BuildCache()

	return &Dependencies{
		WriteDB:     writeDb,
		ReadDB:      readDb,
		CacheClient: client,
	}, nil
}
