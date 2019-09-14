package main

import (
	"api/routes"
	"log"
	"mux"
	"net/http"
)

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/api/payment/transaction", routes.ListAllTransactions)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
