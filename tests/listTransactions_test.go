package listTransactions

import (
	"api/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListAllTransactions(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/payment/transaction", nil)
	if err != nil {
		t.Fatal(err)
	}

	requestRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.ListAllTransactions)
	handler.ServeHTTP(requestRecorder, req)

	if status := requestRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
