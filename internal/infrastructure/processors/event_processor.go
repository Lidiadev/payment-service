package processors

import "payment-service/internal/domain/events"

type EventProcessor interface {
	Process(event events.Event) error
}
