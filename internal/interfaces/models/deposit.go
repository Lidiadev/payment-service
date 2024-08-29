package models

type DepositRequest struct {
	GatewayID  string            `json:"gateway_id" validate:"required,min=1"`
	CustomerID string            `json:"customer_id" validate:"required,min=1"`
	Amount     float64           `json:"amount" validate:"gte=0"`
	Details    map[string]string `json:"details"`
}

type DepositResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
}
