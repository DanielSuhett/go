package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"products/models"
)

func AllProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	w.Header().Set("Content-type", "application/json")

	data, err := json.Marshal(products)

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	fmt.Println(data)

	w.Write(data)
}
