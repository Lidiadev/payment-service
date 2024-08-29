package gateways

import (
	"fmt"
	"payment-service/internal/domain"
)

type GatewayFactoryInterface interface {
	RegisterGateway(name domain.GatewayName, strategy domain.PaymentGatewayStrategy)
	GetGateway(name string) (domain.PaymentGatewayStrategy, error)
}

type GatewayFactory struct {
	gateways map[domain.GatewayName]domain.PaymentGatewayStrategy
}

func NewGatewayFactory() *GatewayFactory {
	return &GatewayFactory{
		gateways: make(map[domain.GatewayName]domain.PaymentGatewayStrategy),
	}
}

func (f *GatewayFactory) RegisterGateway(name domain.GatewayName, strategy domain.PaymentGatewayStrategy) {
	f.gateways[name] = strategy
}

func (f *GatewayFactory) GetGateway(name string) (domain.PaymentGatewayStrategy, error) {
	gatewayName := domain.GatewayName(name)

	gateway, exists := f.gateways[gatewayName]
	if !exists {
		return nil, fmt.Errorf("gateway %s not registered", name)
	}
	return gateway, nil
}
