package pkg

import (
	"encoding/json"
	"fmt"
)

type InMemoryQueue struct {
	handlers map[string]func(message interface{}) error
}

func NewQueue() Queue {
	return &InMemoryQueue{
		handlers: make(map[string]func(message interface{}) error),
	}
}

func (q *InMemoryQueue) Publish(topic string, message interface{}) error {
	handler, ok := q.handlers[topic]
	if !ok {
		return fmt.Errorf("no handler registered for topic: %s", topic)
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return handler(messageBytes)
}

func (q *InMemoryQueue) Subscribe(topic string, handler func(message interface{}) error) error {
	q.handlers[topic] = handler
	return nil
}
