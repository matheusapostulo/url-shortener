package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func openConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func BuildDb() (proxysql *sql.DB, err error) {
	proxysqlDsn := "user:pass@tcp(localhost:6033)/urls"

	proxysql, err = openConnection(proxysqlDsn)
	if err != nil {
		return nil, err
	}

	return proxysql, nil
}
