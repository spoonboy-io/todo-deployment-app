package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Connect creates connection to Postgres DB
func Connect(host, port, user, pwd, database string) (*sql.DB, error ){

	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, database)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		return db, err
	}

	// ensure we fail if connection bad
	if err := db.Ping(); err != nil {
		return db, err
	}

	return db, nil
}