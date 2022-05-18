package main

import (
	"gorm/database"
	"gorm/routes"
	"log"
	"net/http"
)

func main() {
	database.Connect()
	log.Fatal(http.ListenAndServe(":8000", routes.Router()))
}
