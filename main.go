package main

import (
	"log"
	"mux"
	"net/http"
	"api/routes"
)

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/api/payment/transaction", routes.ListAllTransactions)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}