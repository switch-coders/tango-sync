package rabbitmq

import (
	"github.com/streadway/amqp"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
)

func Connect() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqps://ygnhdzkd:DiwirY09UxYt7hzzbEPBWnpoNLCUyCSn@hornet.rmq.cloudamqp.com/ygnhdzkd")
	if err != nil {
		return nil, errors.NewInternalServerError(errors.ErrorConnectingAMQP.GetMessageWithParams(errors.Parameters{"cause": err.Error()}))
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.NewInternalServerError(errors.ErrorOpeningChannel.GetMessageWithParams(errors.Parameters{"cause": err.Error()}))
	}

	return ch, nil
}
