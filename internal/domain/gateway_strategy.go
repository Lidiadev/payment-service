package domain

import "context"

type PaymentGatewayStrategy interface {
	Deposit(ctx context.Context, amount float64, details map[string]string) (string, error)
	Withdraw(ctx context.Context, amount float64, details map[string]string) (string, error)
}
