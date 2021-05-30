package mock

import "github.com/streadway/amqp"

type Channel struct {
}

func (c Channel) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return nil
}
