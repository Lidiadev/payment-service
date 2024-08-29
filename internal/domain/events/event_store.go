package events

type EventStore interface {
	SaveEvent(event Event) error
	GetEvents() ([]Event, error)
}
