package processors

import (
	"fmt"
	"payment-service/internal/domain"
	"payment-service/internal/domain/events"
	"payment-service/internal/infrastructure/gateways"
)

type EventProcessorFactoryInteface interface {
	GetProcessor(eventName string) (EventProcessor, error)
}

type EventProcessorFactory struct {
	eventStore     events.EventStore
	gatewayFactory gateways.GatewayFactoryInterface
}

func NewEventProcessorFactory(eventStore events.EventStore, gatewayFactory gateways.GatewayFactoryInterface) *EventProcessorFactory {
	return &EventProcessorFactory{
		eventStore:     eventStore,
		gatewayFactory: gatewayFactory,
	}
}

func (f *EventProcessorFactory) GetProcessor(eventName string) (EventProcessor, error) {
	switch eventName {
	case domain.DepositReceivedEvent:
		return NewDepositReceivedProcessor(f.eventStore, f.gatewayFactory), nil
	// Other events will be added here
	default:
		return nil, fmt.Errorf("unknown event type: %s", eventName)
	}
}
