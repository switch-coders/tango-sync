package rabbitmq

import (
	"os"
	"strings"

	"github.com/streadway/amqp"

	"github.com/switch-coders/tango-sync/src/api/config/rabbitmq/mock"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/providers"
)

func Connect() (providers.Channel, error) {
	var (
		conn *amqp.Connection
		err  error
	)

	scope := os.Getenv("SCOPE")
	if strings.HasSuffix(scope, "master") || strings.HasSuffix(scope, "beta") {
		if strings.HasSuffix(scope, "master") {
			conn, err = amqp.Dial("amqps://dkjhcutt:G4mvE9HD9-_p4xxoRm74tSMU8EJctEMO@eagle.rmq.cloudamqp.com/dkjhcutt")
		}

		if strings.HasSuffix(scope, "beta") {
			conn, err = amqp.Dial("amqps://zzstsozt:4HMR-rOwRRidMwRmz23Mej8umOuBwzTO@eagle.rmq.cloudamqp.com/zzstsozt")
		}

		if err != nil {
			return nil, errors.NewInternalServerError(errors.ErrorConnectingAMQP.GetMessageWithParams(errors.Parameters{"cause": err.Error()}))
		}

		ch, err := conn.Channel()
		if err != nil {
			return nil, errors.NewInternalServerError(errors.ErrorOpeningChannel.GetMessageWithParams(errors.Parameters{"cause": err.Error()}))
		}

		return ch, nil
	}

	mockCh := new(mock.Channel)
	return mockCh, nil
}
