package models

type PaymentProviderA struct {
	Transactions []FlyPayA `json:"transactions"`
}
type FlyPayA struct {
	Amount         int32  `json:"amount"`
	Currency       string `json:"currency"`
	StatusCode     int    `json:"statusCode"`
	OrderReference string `json:"orderReference"`
	TransactionID  string `json:"transactionId"`
}
type PaymentProviderB struct {
	Transactions []FlyPayB `json:"transactions"`
}
type FlyPayB struct {
	Value               int32  `json:"value"`
	TransactionCurrency string `json:"transactionCurrency"`
	StatusCode          int    `json:"statusCode"`
	OrderInfo           string `json:"orderInfo"`
	PaymentID           string `json:"paymentId"`
}

var StatusCodes []struct {
	Status string
	Codes  []int
}

// return obj provide all current supported payment provider
func SuportedPaymentProviders() map[string]interface{} {
	supportedPaymentProviders := make(map[string]interface{})
	supportedPaymentProviders["flyPayA"] = PaymentProviderA{Transactions: []FlyPayA{}}
	supportedPaymentProviders["flyPayB"] = PaymentProviderB{Transactions: []FlyPayB{}}

	return supportedPaymentProviders
}
