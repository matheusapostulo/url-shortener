package dependencies

import (
	"database/sql"

	"github.com/matheusapostulo/url-shortener/cmd/url/config"
	"github.com/matheusapostulo/url-shortener/internal/url/infra/connection"
	"github.com/redis/go-redis/v9"
)

type Dependencies struct {
	DB           *sql.DB
	CacheClient  *redis.Client
	RabbitMQConn *connection.RabbitMQConnection
}

func BuildDependencies() (*Dependencies, error) {
	//db
	proxysql, err := config.BuildDb()
	if err != nil {
		return nil, err
	}

	// cache
	client := config.BuildCache()

	// rabbitmq
	conn := connection.NewRabbitMQConnection()
	return &Dependencies{
		DB:           proxysql,
		CacheClient:  client,
		RabbitMQConn: conn,
	}, nil

}
