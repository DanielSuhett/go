package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gorm/database"
	"gorm/models"
	"io/ioutil"
	"net/http"
)

func AllProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var products []models.Product

	result := database.DB.Find(&products)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(result.Error.Error()))
	}

	w.Header().Set("Content-type", "application/json")

	data, _ := json.Marshal(products)

	w.Write(data)
}

func GetProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	var product models.Product

	result := database.DB.Find(&product, id).Limit(1)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(result.Error.Error()))
	}

	w.Header().Set("Content-type", "application/json")

	data, _ := json.Marshal(product)

	w.Write(data)
}

func CreateProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var product models.Product

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	result := database.DB.Create(&product)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(result.Error.Error()))
	}

	w.Header().Set("Content-type", "application/json")

	data, _ := json.Marshal(product)

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var product models.Product
	id := ps.ByName("id")

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	result := database.DB.Model(models.Product{}).Where("id = ?", id).Updates(&product)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(result.Error.Error()))
	}

	w.Header().Set("Content-type", "application/json")

	data, _ := json.Marshal(product)

	w.Write(data)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result := database.DB.Delete(models.Product{}, id)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(result.Error.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
