package events

import (
	"github.com/google/uuid"
	"time"
)

type (
	EventPayload interface{}

	Event struct {
		Entity
		Payload    EventPayload
		OccurredAt time.Time
	}
)

func (e Event) EventName() string { return e.Name }

func NewEvent(name string, payload EventPayload) Event {
	return newEvent2(name, payload)
}

func newEvent2(name string, payload EventPayload) Event {
	evt := Event{
		Entity:     NewEntity(uuid.New().String(), name),
		Payload:    payload,
		OccurredAt: time.Now(),
	}

	return evt
}
