package gateways

import (
	"context"
)

type GatewayA struct {
	client GatewayAClientInterface
}

func NewGatewayA(client GatewayAClientInterface) *GatewayA {
	return &GatewayA{
		client: client,
	}
}

func (g *GatewayA) Deposit(ctx context.Context, amount float64, details map[string]string) (string, error) {
	responseData, err := g.client.Deposit(ctx, amount, details)
	if err != nil {
		return "", err
	}

	return responseData, nil
}

func (g *GatewayA) Withdraw(ctx context.Context, amount float64, details map[string]string) (string, error) {
	responseData, err := g.client.Withdraw(ctx, amount, details)
	if err != nil {
		return "", err
	}

	return responseData, nil
}
