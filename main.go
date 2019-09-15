package main

import (
	"api/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.HandleRoutes()
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
