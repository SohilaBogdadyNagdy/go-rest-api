package main

import (
	"api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mux"
	"net/http"
)

func readJSONFile(filePath string) (content []byte) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return content
}

func listAllTransactions(writer http.ResponseWriter, request *http.Request) {
	flyPayAContent := readJSONFile("data/flyPayA.json")
	flyPayBContent := readJSONFile("data/flyPayB.json")

	var FlyPayAResonse models.PaymentProviderA
	err2 := json.Unmarshal(flyPayAContent, &FlyPayAResonse)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	var FlyPayBResonse models.PaymentProviderB
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
