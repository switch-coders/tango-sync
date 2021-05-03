package sync_by_product

import (
	"context"
	goErrors "errors"

	"github.com/getsentry/sentry-go"
	"github.com/switch-coders/tango-sync/src/api/core/contracts/sync_by_product"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/providers"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Implementation struct {
	TNProvider      providers.TiendaNube
	ProductProvider providers.Product
}

func (uc *Implementation) Execute(ctx context.Context, request sync_by_product.Request) error {
	ctx = context.WithValue(ctx, infrastructure.ActionKey, "sync_by_product")

	f := filter.SearchProduct{}
	f.Fields = []string{"variants"}

	f.Q = request.SKU
	variant, err := uc.TNProvider.SearchStock(ctx, f)
	if err != nil {
		if goErrors.As(err, &errors.NotFoundError{}) {
			return nil
		}
	} else {
		if variant == nil {
			return nil
		}

		if !stockIsEqual(int(request.Stock), variant.Stock) {
			variant.Stock = int(request.Stock)
		}
	}

	err = uc.TNProvider.UpdateProductVariant(ctx, *variant)
	if err != nil {
		sentry.CaptureMessage(err.Error())
	} else {
		err = uc.ProductProvider.UpdateOrCreate(ctx, *variant)
		if err != nil {
			sentry.CaptureMessage(err.Error())
		}
	}

	return nil
}

func stockIsEqual(q1, q2 int) bool {
	return q1 == q2
}
