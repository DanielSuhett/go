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

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")

	product := models.GetProduct(id)
	
	w.Header().Set("Content-type", "application/json")

	data, err := json.Marshal(product)

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	fmt.Println(data)

	w.Write(data)
}


func CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	product := models.CreateProduct(body)
	
	w.Header().Set("Content-type", "application/json")

	data, err := json.Marshal(product)

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(201)
	w.Write(data)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	product := models.UpdateProduct(body)
	
	w.Header().Set("Content-type", "application/json")

	data, err := json.Marshal(product)

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(201)
	w.Write(data)
}