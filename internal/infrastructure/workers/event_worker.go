package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"payment-service/internal/domain"
	"payment-service/internal/domain/events"
	"payment-service/internal/infrastructure/gateways"
	"payment-service/pkg"
)

type EventWorker struct {
	queue           pkg.Queue
	deadLetterQueue pkg.Queue
	eventStore      events.EventStore
	gatewayFactory  gateways.GatewayFactoryInterface
}

func NewEventWorker(queue pkg.Queue, deadLetterQueue pkg.Queue, eventStore events.EventStore, gatewayFactory gateways.GatewayFactoryInterface) *EventWorker {
	return &EventWorker{
		queue:           queue,
		deadLetterQueue: deadLetterQueue,
		eventStore:      eventStore,
		gatewayFactory:  gatewayFactory,
	}
}

func (w *EventWorker) Start(ctx context.Context) error {
	for _, eventType := range []string{domain.DepositReceivedEvent /*domain.WithdrawRequestedEvent*/} {
		if err := w.queue.Subscribe(eventType, func(message interface{}) error {
			return w.processMessage(message.([]byte))
		}); err != nil {
			return fmt.Errorf("failed to subscribe to %s: %w", eventType, err)
		}
	}

	<-ctx.Done()
	return nil
}

func (w *EventWorker) processMessage(message []byte) error {
	var event events.Event
	if err := json.Unmarshal(message, &event); err != nil {
		return fmt.Errorf("failed to unmarshal base event: %w", err)
	}

	switch event.EventName() {
	case domain.DepositReceivedEvent:
		return w.processDepositReceived(event)
	//case domain.WithdrawRequestedEvent:
	//	return w.processWithdrawRequested(message)
	default:
		return fmt.Errorf("unknown event type: %s", event.EventName())
	}
}

func (w *EventWorker) processDepositReceived(event events.Event) error {
	payloadBytes, err := json.Marshal(event.Payload)
	if err != nil {
		return fmt.Errorf("failed to marshal event payload: %w", err)
	}

	var depositReceived domain.DepositReceived
	if err := json.Unmarshal(payloadBytes, &depositReceived); err != nil {
		return fmt.Errorf("failed to unmarshal DepositReceived event: %w", err)
	}

	gateway, err := w.gatewayFactory.GetGateway(depositReceived.Gateway)
	if err != nil {
		return w.handleFailure(event, fmt.Errorf("unknown gateway: %s", depositReceived.Gateway))
	}

	_, err = gateway.Deposit(context.Background(), depositReceived.Amount, nil)

	if err != nil {
		return w.handleFailure(event, fmt.Errorf("failed to process deposit: %w", err))
	}

	depositProcessed := domain.DepositProcessed{
		TransactionID: depositReceived.TransactionID,
		Status:        domain.ProcessedStatus,
		Gateway:       depositReceived.Gateway,
	}
	processedEvent := events.NewEvent2(domain.DepositProcessedEvent, &depositProcessed)

	return w.eventStore.SaveEvent(processedEvent)
}

func (w *EventWorker) handleFailure(baseEvent events.Event, err error) error {
	//var failedEvent events.Event

	log.Printf("Transaction failed: %s %v", baseEvent.Id, err)

	err = w.deadLetterQueue.Publish(baseEvent.EventName(), baseEvent)
	if err != nil {
		return err
	}

	//switch baseEvent.EventName() {
	//case domain.DepositReceivedEvent:
	//	depositReceived := domain.DepositFailed{
	//		TransactionID: uuid.New().String(),
	//		Reason:        err.Error(),
	//	}
	//	failedEvent = events.NewEvent2(domain.DepositFailedEvent, &depositReceived)
	//	//failedEvent = domain.DepositFailed{
	//	//	BaseEvent: domain.BaseEvent{
	//	//		EventID:     generateUUID(),
	//	//		AggregateId: baseEvent.AggregateID(),
	//	//		OccurredAt:  time.Now(),
	//	//		EventType:   domain.DepositFailedEvent,
	//	//	},
	//	//	Reason: err.Error(),
	//	//}
	////case domain.WithdrawRequestedEvent:
	////	failedEvent = domain.WithdrawFailed{
	////		BaseEvent: domain.BaseEvent{
	////			EventID:     generateUUID(),
	////			AggregateId: baseEvent.AggregateID(),
	////			OccurredAt:  time.Now(),
	////			EventType:   domain.WithdrawFailedEvent,
	////		},
	////		Reason: err.Error(),
	////	}
	//default:
	//	return fmt.Errorf("unknown event type for failure handling: %s", baseEvent.EventName())
	//}
	//
	//if err := w.eventStore.SaveEvent(failedEvent); err != nil {
	//	return fmt.Errorf("failed to save failure event: %w", err)
	//}
	//
	//log.Printf("Transaction failed: %v", err)
	//return nil
	return nil
}
