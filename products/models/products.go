package models

import (
	db "products/database"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	rows, err := db.Query("select * from Products")

	if err != nil {
		panic(err)
	}

	products := []Product{}

	for rows.Next() {
		var name string
		var description string
		var price float64
		var quantity int

		err := rows.Scan(&name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}

		products = append(products, Product{name, description, price, quantity})
	}

	return products
}