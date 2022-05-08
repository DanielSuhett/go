package models

func GetAllProducts () {
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
}