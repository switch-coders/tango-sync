package account

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/switch-coders/tango-sync/src/api/core/contracts/integration"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Repository struct {
	DBClient *gorm.DB
}

func (r *Repository) Create(ctx context.Context, req integration.Request) error {
	transaction := newrelic.FromContext(ctx)

	entity := newAccountEntity(req)

	var err error

	infrastructure.WrapDatastoreSegment("Postgres", "SAVE-ACCOUNT", transaction, func() {
		err = r.DBClient.Save(&entity).Error
	})

	if err != nil {
		sentry.CaptureException(err)

		return errors.NewRepositoryError(errors.ErrorUpsertResource.GetMessageWithParams(errors.Parameters{"resource": "save_account", "name": req.Name}))
	}

	return nil
}
