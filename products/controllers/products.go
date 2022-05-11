package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"products/models"

	"github.com/julienschmidt/httprouter"
)


func AllProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func GetProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

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


func CreateProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func UpdateProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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