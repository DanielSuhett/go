package models

import (
	"fmt"
	db "products/database"
)

type Product struct {
	id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	rows, err := db.Query("*", "")

	if err != nil {
		panic(err)
	}

	products := []Product{}

	for rows.Next() {
		var id int
		var name string
		var description string
		var price float64
		var quantity int

		err := rows.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}

		products = append(products, Product{id, name, description, price, quantity})
	}

	return products
}

func GetProduct(id string) Product {
	where := fmt.Sprintf("where id = %v", id)

	rows, err := db.Query("*", where)

	if err != nil {
		panic(err)
	}

	product := Product{}

	for rows.Next() {
		var id int
		var name string
		var description string
		var price float64
		var quantity int

		err := rows.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}

		product = Product{id, name, description, price, quantity}
	}

	return product
}

func CreateProduct(values db.Product) {
	product := db.Product{Name: values.Name, Description: values.Description, Price: values.Price, Quantity: values.Quantity}

	_, err := db.Insert("name, description, price, quantity", product)

	if err != nil {
		panic(err)
	}
}

func UpdateProduct(data ...interface{}) Product {
	product := Product{}

	return product
}
