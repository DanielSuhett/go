package main

import (
	"net/http"
	db "products/database"
	"text/template"
)



type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var view = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}


func index(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from Products")

	if err != nil {
		panic(err)
	}

	products := []Product{}

	for rows.Next() {
		var name        string
		var	description string
		var	price       float64
		var	quantity    int
		

		err := rows.Scan(&name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}


		products = append(products, Product{name, description, price, quantity})
	}

	view.ExecuteTemplate(w, "index", products)
}
