package gateways

import (
	"context"
)

type GatewayB struct {
	client GatewayBClientInterface
}

func NewGatewayB(client GatewayBClientInterface) *GatewayB {
	return &GatewayB{
		client: client,
	}
}

func (g *GatewayB) Deposit(ctx context.Context, amount float64, details map[string]string) (string, error) {
	responseData, err := g.client.Deposit(ctx, amount, details)
	if err != nil {
		return "", err
	}

	return responseData, nil
}

func (g *GatewayB) Withdraw(ctx context.Context, amount float64, details map[string]string) (string, error) {
	responseData, err := g.client.Withdraw(ctx, amount, details)
	if err != nil {
		return "", err
	}

	return responseData, nil
}
