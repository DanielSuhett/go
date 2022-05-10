package models

import (
	"fmt"
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

func GetProduct(id string) Product {
    query := fmt.Sprintf("select * from Products where id = %v", id)
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	product := Product{}

	for rows.Next() {
		var name string
		var description string
		var price float64
		var quantity int

		err := rows.Scan(&name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}

		product = Product{name, description, price, quantity}
	}

	return product
}


func CreateProduct(data ...interface{}) Product {
	// insert no banco

	// if err != nil {
	// 	panic(err)
	// }

	product := Product{}

	// for rows.Next() {
	// 	var name string
	// 	var description string
	// 	var price float64
	// 	var quantity int

	// 	err := rows.Scan(&name, &description, &price, &quantity)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	product = Product{name, description, price, quantity}
	// }

	return product
}

func UpdateProduct(data ...interface{}) Product {
	// insert no banco

	// if err != nil {
	// 	panic(err)
	// }

	product := Product{}

	// for rows.Next() {
	// 	var name string
	// 	var description string
	// 	var price float64
	// 	var quantity int

	// 	err := rows.Scan(&name, &description, &price, &quantity)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	product = Product{name, description, price, quantity}
	// }

	return product
}