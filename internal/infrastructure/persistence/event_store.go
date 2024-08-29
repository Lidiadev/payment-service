package persistence

import (
	"payment-service/internal/domain/events"
	"payment-service/internal/infrastructure/logging"
	"sync"
)

type InMemoryEventStore struct {
	mu     sync.Mutex
	events []events.Event
	logger logging.LoggerInterface
}

func NewInMemoryEventStore(logger logging.LoggerInterface) *InMemoryEventStore {
	return &InMemoryEventStore{
		logger: logger,
	}
}

func (s *InMemoryEventStore) SaveEvent(event events.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)

	//s.logger.Info("Event saved: " + string(events.Type))

	return nil
}

func (s *InMemoryEventStore) GetEvents() ([]events.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.events, nil
}
