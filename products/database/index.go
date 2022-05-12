package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

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

func Query(fields string, where string) (*sql.Rows, error) {
	db := Connect()
	rows, err := db.Query(fmt.Sprintf("select %v from Products %v", fields, where))

	if err != nil {
		return nil, error(err)
	}

	defer db.Close()

	return rows, nil
}

func Insert(fields string, values interface{}) (sql.Result, error) {
	db := Connect()

	insert, err := db.Prepare(fmt.Sprintf("insert into Products(%v) values(?, ?, ?, ?)", fields))

	if err != nil {
		return nil, error(err)
	}

	product := values.(Product)

	result, err := insert.Exec(product.Name, product.Description, product.Price, product.Quantity)

	if err != nil {
		return nil, error(err)
	}

	defer db.Close()

	return result, nil
}
