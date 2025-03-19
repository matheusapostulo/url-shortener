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

func BuildDb() (writeDB *sql.DB, readDB *sql.DB, err error) {
	master := "root:root@tcp(localhost:3306)/urls"
	mySqlRouter := "root:root@tcp(localhost:3307)/urls"

	writeDB, err = openConnection(master)
	if err != nil {
		return nil, nil, err
	}

	readDB, err = openConnection(mySqlRouter)
	if err != nil {
		return nil, nil, err
	}

	return writeDB, readDB, nil
}
