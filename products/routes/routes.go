package routes

import (
	"net/http"
	"products/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.CreateProduct)
	http.HandleFunc("/:id", controllers.UpdateProduct)
	http.HandleFunc("/products", controllers.AllProducts)
	http.HandleFunc("/product/:id", controllers.GetProduct)
}