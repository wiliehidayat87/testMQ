package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitPSQL(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db, nil
}
