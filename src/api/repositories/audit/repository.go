package audit

import (
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Repository struct {
	DBClient *gorm.DB
}

func (r *Repository) Create(ctx context.Context, sku, j string, v interface{}) error {
	transaction := newrelic.FromContext(ctx)

	entity := new(audit)
	entity.Sku = sku
	entity.Job = j
	entity.Value = fmt.Sprintf("%v", v)

	var err error

	infrastructure.WrapDatastoreSegment("Postgres", "SAVE-AUDIT", transaction, func() {
		err = r.DBClient.Save(&entity).Error
	})

	if err != nil {
		sentry.CaptureException(err)

		return errors.NewRepositoryError(errors.ErrorUpsertResource.GetMessageWithParams(errors.Parameters{"resource": "save_audit", "sku": sku}))
	}

	return nil
}
