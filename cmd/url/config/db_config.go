package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLCluster struct {
	master string
	slaves []string
}

func NewMySQLCluster(master string, slaves []string) *MySQLCluster {
	return &MySQLCluster{
		master: master,
		slaves: slaves,
	}
}

func getActiveThreads(db *sql.DB) (int, error) {
	rows, err := db.Query("SHOW PROCESSLIST")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var activeThreads int
	for rows.Next() {
		var id, user, host, command, state, info any
		var time int
		var dbName sql.NullString

		if err := rows.Scan(&id, &user, &host, &dbName, &command, &time, &state, &info); err != nil {
			return 0, err
		}

		if command != "Sleep" {
			activeThreads++
		}
	}

	if err := rows.Err(); err != nil {
		return 0, err
	}

	return activeThreads, nil
}

func (cluster *MySQLCluster) getLeastLoadedSlave() (*sql.DB, error) {
	var leastLoadedSlave *sql.DB
	leastThreads := int(^uint(0) >> 1)

	for _, slave := range cluster.slaves {
		db, err := sql.Open("mysql", slave)
		if err != nil {
			log.Printf("Error connecting to slave %s: %v\n", slave, err)
			continue
		}

		activeThreads, err := getActiveThreads(db)
		if err != nil {
			log.Printf("Error getting active threads for slave %s: %v\n", slave, err)
			db.Close()
			continue
		}

		if activeThreads < leastThreads {
			leastThreads = activeThreads
			leastLoadedSlave = db
		} else {
			db.Close()
		}
	}

	if leastLoadedSlave == nil {
		return nil, fmt.Errorf("no slaves available")
	}

	return leastLoadedSlave, nil
}

func (cluster *MySQLCluster) getWriteConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", cluster.master)
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
	slaves := []string{
		"root:root@tcp(localhost:3307)/urls",
		"root:root@tcp(localhost:3308)/urls",
	}

	cluster := NewMySQLCluster(master, slaves)

	writeDB, err = cluster.getWriteConnection()
	if err != nil {
		return nil, nil, err
	}

	readDB, err = cluster.getLeastLoadedSlave()
	if err != nil {
		return nil, nil, err
	}

	return writeDB, readDB, nil
}
