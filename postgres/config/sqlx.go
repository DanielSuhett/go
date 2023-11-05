package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetReaderSqlx() *sqlx.DB {
	reader := sqlx.MustConnect("postgres", DB_CONNECTION)

	return reader
}

func GetWriterSqlx() *sqlx.DB {
	writer := sqlx.MustConnect("postgres", DB_CONNECTION)

	return writer
}
