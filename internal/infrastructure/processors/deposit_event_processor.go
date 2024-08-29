package processors

import (
	"context"
	"encoding/json"
	"fmt"
	"payment-service/internal/domain"
	"payment-service/internal/domain/events"
	"payment-service/internal/infrastructure/gateways"
)

type DepositReceivedProcessor struct {
	eventStore     events.EventStore
	gatewayFactory gateways.GatewayFactoryInterface
}

func NewDepositReceivedProcessor(eventStore events.EventStore, gatewayFactory gateways.GatewayFactoryInterface) *DepositReceivedProcessor {
	return &DepositReceivedProcessor{
		eventStore:     eventStore,
		gatewayFactory: gatewayFactory,
	}
}

func (p *DepositReceivedProcessor) Process(event events.Event) error {
	payloadBytes, err := json.Marshal(event.Payload)
	if err != nil {
		return fmt.Errorf("failed to marshal event payload: %w", err)
	}

	var depositReceived domain.DepositReceived
	if err := json.Unmarshal(payloadBytes, &depositReceived); err != nil {
		return fmt.Errorf("failed to unmarshal DepositReceived event: %w", err)
	}

	gateway, err := p.gatewayFactory.GetGateway(depositReceived.Gateway)
	if err != nil {
		return fmt.Errorf("unknown gateway: %s", depositReceived.Gateway)
	}

	_, err = gateway.Deposit(context.Background(), depositReceived.Amount, nil)
	if err != nil {
		return fmt.Errorf("failed to process deposit: %w", err)
	}

	depositProcessed := domain.DepositProcessed{
		TransactionID: depositReceived.TransactionID,
		Status:        domain.ProcessedStatus,
		Gateway:       depositReceived.Gateway,
	}
	processedEvent := events.NewEvent(domain.DepositProcessedEvent, &depositProcessed)

	return p.eventStore.SaveEvent(processedEvent)
}
