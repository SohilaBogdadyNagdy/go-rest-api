package routes

import (
	"api/helpers"
	"api/models"
	"encoding/json"

	// "fmt"

	"net/http"
)

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
		response = helpers.ParseJSONFile("data/"+byProviderFilter+".json", providerExist)

	} else {
		var allTrasn []interface{}
		for providerName, providerStruct := range models.SuportedPaymentProviders() {
			paymentTrans := helpers.ParseJSONFile("data/"+providerName+".json", providerStruct)
			allTrasn = append(allTrasn, paymentTrans)

		}
		response = allTrasn
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
