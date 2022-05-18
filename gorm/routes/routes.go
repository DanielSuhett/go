package routes

import (
	"github.com/julienschmidt/httprouter"
	"gorm/controllers"
)

func Router() *httprouter.Router {
	router := httprouter.New()
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products", controllers.AllProducts)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.GET("/products/:id", controllers.GetProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)

	return router
}
