package events

import (
	"github.com/google/uuid"
	"time"
)

type (
	EventPayload interface{}

	//Event interface {
	//	IDer
	//	EventName() string
	//	Payload() EventPayload
	//	//Metadata() Metadata
	//	OccurredAt() time.Time
	//}

	Event struct {
		Entity
		Payload EventPayload
		//metadata   Metadata
		OccurredAt time.Time
	}

	BaseEvent struct {
		Entity
		Payload EventPayload
		//metadata   Metadata
		OccurredAt time.Time
	}
)

//var _ Event = (*BaseEvent)(nil)

func NewEvent(name string, payload EventPayload /*options ...EventOption*/) BaseEvent {
	return newEvent(name, payload /*options...*/)
}

func newEvent(name string, payload EventPayload /*options ...EventOption*/) BaseEvent {
	evt := BaseEvent{
		Entity:  NewEntity(uuid.New().String(), name),
		Payload: payload,
		//metadata:   make(Metadata),
		OccurredAt: time.Now(),
	}

	//for _, option := range options {
	//	option.configureEvent(&evt)
	//}

	return evt
}

func (e Event) EventName() string { return e.Name }

func (e BaseEvent) EventName() string { return e.Name }

//func (e BaseEvent) Payload() EventPayload { return e.Payload }
//
//// func (e BaseEvent) Metadata() Metadata    { return e.metadata }
//func (e BaseEvent) OccurredAt() time.Time { return e.OccurredAt }

func NewEvent2(name string, payload EventPayload /*options ...EventOption*/) Event {
	return newEvent2(name, payload /*options...*/)
}

func newEvent2(name string, payload EventPayload) Event {
	evt := Event{
		Entity:     NewEntity(uuid.New().String(), name),
		Payload:    payload,
		OccurredAt: time.Now(),
	}

	return evt
}
