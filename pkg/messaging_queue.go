package pkg

type Queue interface {
	Publish(topic string, message interface{}) error
	Subscribe(topic string, handler func(message interface{}) error) error
}
