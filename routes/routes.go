package routes

import (
	"api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Read Json file and parsed it
func ParseJSONFile(filePath string, parsedTo interface{}) (ParsedData interface{}) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	err2 := json.Unmarshal(content, &parsedTo)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return parsedTo
}

//List transactions in payment providers
func ListAllTransactions(writer http.ResponseWriter, request *http.Request) {
	var paymentATransactions models.PaymentProviderA
	respA := ParseJSONFile("data/flyPayA.json", paymentATransactions)

	var paymentBTransactions models.PaymentProviderB
	respB := ParseJSONFile("data/flyPayB.json", paymentBTransactions)

	response := []interface{}{respA, respB}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
