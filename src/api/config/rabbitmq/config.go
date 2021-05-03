package rabbitmq

import (
	"github.com/streadway/amqp"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"os"
	"strings"
)

func Connect() (*amqp.Channel, error) {
	conn, err := setupEnvironment()
	if err != nil {
		return nil, errors.NewInternalServerError(errors.ErrorConnectingAMQP.GetMessageWithParams(errors.Parameters{"cause": err.Error()}))
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.NewInternalServerError(errors.ErrorOpeningChannel.GetMessageWithParams(errors.Parameters{"cause": err.Error()}))
	}

	return ch, nil
}

func setupEnvironment() (*amqp.Connection, error) {
	scope := os.Getenv("SCOPE")
	if strings.HasSuffix(scope, "beta") {
		return amqp.Dial("amqps://zzstsozt:4HMR-rOwRRidMwRmz23Mej8umOuBwzTO@eagle.rmq.cloudamqp.com/zzstsozt")
	}

	return amqp.Dial("amqps://zzstsozt:4HMR-rOwRRidMwRmz23Mej8umOuBwzTO@eagle.rmq.cloudamqp.com/zzstsozt")
}
