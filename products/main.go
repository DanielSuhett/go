package main

import (
	"log"
	"net/http"
	"products/routes"
)

func main() {
	routes.LoadRoutes()
    log.Fatal(http.ListenAndServe(":8000", nil))
}
