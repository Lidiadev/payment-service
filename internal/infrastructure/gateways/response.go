package gateways

type DepositGatewayAResponse struct {
	TransactionID string `json:"transaction_id" xml:"transaction_id"`
	Status        string `json:"status" xml:"status"`
}

type WithdrawalGatewayAResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
}

type DepositGatewayBResponse struct {
	TransactionID string `xml:"transaction_id"`
	Status        string `xml:"status"`
}

type WithdrawalGatewayBResponse struct {
	TransactionID string `xml:"transaction_id"`
	Status        string `xml:"status"`
}

type TransactionIDResponse struct {
	TransactionID string `xml:"transaction_id"`
}
