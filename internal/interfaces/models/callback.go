package models

type CallbackRequest struct {
	TransactionID string                 `json:"transaction_id" validate:"required,min=1"`
	Status        string                 `json:"status" validate:"required,min=1"`
	Details       map[string]interface{} `json:"details"`
}

type GatewayACallbackRequest struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	ExtraData     string `json:"extra_data"`
}

type GatewayBCallbackRequest struct {
	TransactionID string                 `json:"transaction_id" validate:"required,min=1"`
	Status        string                 `json:"status" validate:"required,min=1"`
	Metadata      map[string]interface{} `json:"metadata"`
}

type CallbackResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
}
