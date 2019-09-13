package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mux"
	"net/http"
)

type TransactionsFlyPayA struct {
	Transactions []FlyPayA `json:"transactions"`
}
type FlyPayA struct {
	Amount         int32  `json:"amount"`
	Currency       string `json:"currency"`
	StatusCode     int    `json:"statusCode"`
	OrderReference string `json:"orderReference"`
	TransactionID  string `json:"transactionId"`
}
type TransactionsFlyPayB struct {
	Transactions []FlyPayB `json:"transactions"`
}
type FlyPayB struct {
	Value               int32  `json:"value"`
	TransactionCurrency string `json:"transactionCurrency"`
	StatusCode          int    `json:"statusCode"`
	OrderInfo           string `json:"orderInfo"`
	PaymentID           string `json:"paymentId"`
}

func readJSONFile(filePath string) (content []byte) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return content
}

func listAllTransactions(writer http.ResponseWriter, request *http.Request) {
	flyPayAContent := readJSONFile("paymentProviders/flyPayA.json")
	flyPayBContent := readJSONFile("paymentProviders/flyPayB.json")

	var FlyPayAResonse TransactionsFlyPayA
	err2 := json.Unmarshal(flyPayAContent, &FlyPayAResonse)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	var FlyPayBResonse TransactionsFlyPayB
	err3 := json.Unmarshal(flyPayBContent, &FlyPayBResonse)
	if err3 != nil {
		fmt.Println(err2.Error())
	}
	response := []interface{}{FlyPayAResonse, FlyPayBResonse}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/api/payment/transaction", listAllTransactions)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
