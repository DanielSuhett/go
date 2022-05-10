package routes

import (
	"net/http"
	"products/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/products", controllers.AllProducts)

}