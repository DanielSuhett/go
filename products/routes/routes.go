package routes

import (
	"products/controllers"
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/:id", controllers.UpdateProduct)
    router.GET("/products", controllers.AllProducts)
	router.GET("/products/:id", controllers.GetProduct)

	return router
}