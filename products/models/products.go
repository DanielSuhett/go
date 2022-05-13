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

func GetAllProducts() ([]Product, error) {
	rows, err := db.Query("*", "")

	if err != nil {
		return nil, error(err)
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
			return nil, error(err)
		}

		products = append(products, Product{id, name, description, price, quantity})
	}

	return products, nil
}

func GetProduct(id string) (*Product, error) {
	where := fmt.Sprintf("where id = %v", id)

	rows, err := db.Query("*", where)

	if err != nil {
		return nil, error(err)
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
			return nil, error(err)
		}

		product = Product{id, name, description, price, quantity}
	}

	return &product, nil
}

func CreateProduct(values db.Product) error {
	product := db.Product{Name: values.Name, Description: values.Description, Price: values.Price, Quantity: values.Quantity}

	_, err := db.Insert("name, description, price, quantity", product)

	if err != nil {
		return error(err)
	}

	return nil
}

func DeleteProduct(id int) (error) {
	err := db.Delete(id)

	if err != nil {
		return error(err)
	}

	return nil
}

func UpdateProduct(data ...interface{}) Product {
	product := Product{}

	return product
}
