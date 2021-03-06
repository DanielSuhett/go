package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"products/database"
	"products/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)


func AllProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	products, err := models.GetAllProducts()

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-type", "application/json")

	data, _ := json.Marshal(products)

	w.Write(data)
}

func GetProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	product, err := models.GetProduct(id)

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	
	w.Header().Set("Content-type", "application/json")

	data, _ := json.Marshal(product)

	w.Write(data)
}


func CreateProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var product database.Product

	body, _ := ioutil.ReadAll(r.Body)
    err := json.Unmarshal(body, &product)

	
	if(err != nil){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	fail := models.CreateProduct(product)

	if(fail != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fail.Error()))
	}
	
	w.Header().Set("Content-type", "application/json")

	data, _ := json.Marshal(product)

	w.WriteHeader(201)
	w.Write(data)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var product database.Product
	id := ps.ByName("id")


	body, _ := ioutil.ReadAll(r.Body)
    err := json.Unmarshal(body, &product)

	
	if(err != nil){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	fail := models.UpdateProduct(id, product)

	if(fail != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fail.Error()))
	}

	w.Header().Set("Content-type", "application/json")

	data, _ := json.Marshal(product)

	w.WriteHeader(200)
	w.Write(data)
}


func DeleteProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("id")

	id, _ := strconv.Atoi(param)

	err := models.DeleteProduct(id)
	
	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}