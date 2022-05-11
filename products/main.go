package main

import (
	"log"
	"net/http"
	"products/routes"
)

func main() {
    log.Fatal(http.ListenAndServe(":8000", routes.Router()))
}
