package database

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/GOLANG")

	if err != nil {
		panic(err)
	}

	fmt.Println("Connect with database successfuly")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func Query(request string) (*sql.Rows, error) {
	db := Connect()
	rows, err := db.Query(request)

	if err != nil {
		return nil, error(err)
	}

	defer db.Close()

	return rows, nil
}
