package models

type PaymentProvider interface{}

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
