package product

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Repository struct {
	DBClient *gorm.DB
}

func (r *Repository) UpdateOrCreate(ctx context.Context, p entities.Product) error {
	transaction := newrelic.FromContext(ctx)

	entity := new(product)
	entity.SKU = p.Sku
	entity.Stock = &p.Stock
	entity.Price = &p.Price

	var err error

	infrastructure.WrapDatastoreSegment("Postgres", "CREATE-UPDATE", transaction, func() {
		err = r.DBClient.Save(&entity).Error
	})

	if err != nil {
		sentry.CaptureException(err)

		return errors.NewRepositoryError(errors.ErrorUpsertResource.GetMessageWithParams(errors.Parameters{"resource": "create_or_update_product", "sku": p.Sku}))
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, sku string) (*entities.Product, error) {
	transaction := newrelic.FromContext(ctx)

	var p product
	var err error

	infrastructure.WrapDatastoreSegment("Postgres", "SELECT", transaction, func() {
		err = r.DBClient.Where("sku = ?", sku).First(&p).Error
	})

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.NewNotFoundError(errors.ErrorGettingResource.GetMessageWithParams(errors.Parameters{"resource": "get_product", "sku": sku}))
		}

		sentry.CaptureException(err)
		return nil, errors.NewRepositoryError(errors.ErrorGettingResource.GetMessageWithParams(errors.Parameters{"resource": "get_product", "sku": sku}))
	}

	return p.toEntity(), nil
}
