package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"payment-service/internal/domain"
	"payment-service/internal/domain/events"
	"payment-service/internal/infrastructure/processors"
	"payment-service/pkg"
)

type EventWorker struct {
	queue            pkg.Queue
	deadLetterQueue  pkg.Queue
	processorFactory processors.EventProcessorFactoryInteface
}

func NewEventWorker(queue pkg.Queue, deadLetterQueue pkg.Queue, processorFactory processors.EventProcessorFactoryInteface) *EventWorker {
	return &EventWorker{
		queue:            queue,
		deadLetterQueue:  deadLetterQueue,
		processorFactory: processorFactory,
	}
}

func (w *EventWorker) Start(ctx context.Context) error {
	for _, eventType := range []string{domain.DepositReceivedEvent /* add other events here */} {
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

	processor, err := w.processorFactory.GetProcessor(event.EventName())
	if err != nil {
		return w.handleFailure(event, fmt.Errorf("failed to get processor: %w", err))
	}

	if err := processor.Process(event); err != nil {
		return w.handleFailure(event, fmt.Errorf("failed to process event: %w", err))
	}

	return nil
}

func (w *EventWorker) handleFailure(baseEvent events.Event, err error) error {
	log.Printf("Transaction failed: %s %v", baseEvent.Id, err)

	err = w.deadLetterQueue.Publish(baseEvent.EventName(), baseEvent)
	if err != nil {
		return err
	}

	return nil
}
