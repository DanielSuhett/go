package database

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

func Connect() (*Database, error) {
	connStr := "user=danielsuhett password=daniel dbname=danielsuhett sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

func Close(db *Database) error {
	err := db.Close()
	return err
}
