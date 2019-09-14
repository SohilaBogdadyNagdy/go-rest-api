package routes

import (
	"api/models"
	"encoding/json"

	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Read Json file and parsed it
func ParseJSONFile(filePath string, parsedTo interface{}) (ParsedData interface{}) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err2 := json.Unmarshal(content, &parsedTo)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	return parsedTo
}

//List transactions in payment providers
func ListAllTransactions(writer http.ResponseWriter, request *http.Request) {
	var response interface{}
	query := request.URL.Query()
	byProviderFilter := query.Get("provider")
	//byStatusCodeFilter := query.Get("statusCode")

	if byProviderFilter != "" {
		providerExist := models.SuportedPaymentProviders()[byProviderFilter]
		if providerExist == nil {
			err := map[string]interface{}{"code": "Invalid Provider"}
			writer.Header().Set("Content-type", "applciation/json")
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(err)
			return
		}
		response = ParseJSONFile("data/"+byProviderFilter+".json", providerExist)

	} else {
		var allTrasn []interface{}
		for providerName, providerStruct := range models.SuportedPaymentProviders() {
			paymentTrans := ParseJSONFile("data/"+providerName+".json", providerStruct)
			allTrasn = append(allTrasn, paymentTrans)

		}
		response = allTrasn
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
