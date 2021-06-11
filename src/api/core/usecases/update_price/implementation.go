package update_price

import (
	"context"
	goErrors "errors"

	priceContract "github.com/switch-coders/tango-sync/src/api/core/contracts/sync_by_product/price"
	"github.com/switch-coders/tango-sync/src/api/core/entities/constants"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/providers"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
	"github.com/switch-coders/tango-sync/src/api/util/price"
)

type Implementation struct {
	TNProvider      providers.TiendaNube
	ProductProvider providers.Product
	AuditProvider   providers.Audit
}

func (uc *Implementation) Execute(ctx context.Context, request priceContract.Request) error {
	ctx = context.WithValue(ctx, infrastructure.ActionKey, "update_price")

	f := filter.SearchProduct{}
	f.Fields = []string{"variants"}

	f.Q = request.SKU
	variant, err := uc.TNProvider.SearchProduct(ctx, f)
	if err != nil {
		if goErrors.As(err, &errors.NotFoundError{}) {
			return nil
		}

		return errors.GetError(err, err.Error())
	} else {
		if variant == nil {
			return nil
		}

		if !price.IsEqual(request.Price, variant.Price) {
			variant.Price = request.Price
		}
	}

	err = uc.TNProvider.UpdateProductPriceVariant(ctx, *variant)
	if err == nil {
		err = uc.ProductProvider.UpdateOrCreate(ctx, *variant)
		if err == nil {
			_ = uc.AuditProvider.Create(ctx, variant.Sku, constants.JobUpdatePrice, variant.Price)
		}
	}

	return nil
}
