package commands

import "payment-service/internal/domain"

type ProcessDepositCommand struct {
	CustomerID string
	Amount     float64
	Details    map[string]string
	GatewayID  string
}

type ProcessDepositCommandResult struct {
	TransactionID string
	Status        domain.TransactionStatus
}
