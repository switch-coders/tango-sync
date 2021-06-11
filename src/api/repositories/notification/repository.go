package notification

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/streadway/amqp"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/providers"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Repository struct {
	Channel providers.Channel
}

func (r *Repository) Notify(ctx context.Context, n entities.Notification) error {
	transaction := newrelic.FromContext(ctx)

	msg, _ := json.Marshal(n.Message)

	var err error

	infrastructure.WrapExternalSegmentWithAlias(transaction, n.Topic, "PUBLISH_STOCK", func() {
		err = r.Channel.Publish(
			"",
			n.Topic,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        msg,
			})
	})

	if err != nil {
		sentry.CaptureException(err)

		return errors.NewInternalServerError(errors.ErrorFailsToPublishRabbitMQ.GetMessageWithParams(errors.Parameters{
			"topic": n.Topic,
		}))
	}

	return nil
}
