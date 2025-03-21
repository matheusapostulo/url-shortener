package dependencies

import (
	"database/sql"

	"github.com/matheusapostulo/url-shortener/cmd/url/config"
	"github.com/redis/go-redis/v9"
)

type Dependencies struct {
	DB          *sql.DB
	CacheClient *redis.Client
}

func BuildDependencies() (*Dependencies, error) {
	//db
	proxysql, err := config.BuildDb()
	if err != nil {
		return nil, err
	}

	// cache
	client := config.BuildCache()

	return &Dependencies{
		DB:          proxysql,
		CacheClient: client,
	}, nil
}
