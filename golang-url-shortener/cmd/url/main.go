package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/matheusapostulo/url-shortener/cmd/url/application"
	"github.com/matheusapostulo/url-shortener/cmd/url/dependencies"
)

func main() {
	dependencies, err := dependencies.BuildDependencies()
	if err != nil {
		panic(err)
	}
	defer dependencies.DB.Close()
	defer dependencies.CacheClient.Close()

	application.Run(*dependencies)
}
