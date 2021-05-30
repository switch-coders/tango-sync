package update_stock

import (
	"context"
	goErrors "errors"
	stockContract "github.com/switch-coders/tango-sync/src/api/core/contracts/sync_by_product/stock"
	"github.com/switch-coders/tango-sync/src/api/core/entities/constants"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/providers"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
	"github.com/switch-coders/tango-sync/src/api/util/stock"
)

type Implementation struct {
	TNProvider      providers.TiendaNube
	ProductProvider providers.Product
	AuditProvider   providers.Audit
}

func (uc *Implementation) Execute(ctx context.Context, request stockContract.Request) error {
	ctx = context.WithValue(ctx, infrastructure.ActionKey, "update_stock")

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

		if !stock.IsEqual(int(request.Stock), variant.Stock) {
			variant.Stock = int(request.Stock)
		}
	}

	err = uc.TNProvider.UpdateProductStockVariant(ctx, *variant)
	if err == nil {
		err = uc.ProductProvider.UpdateOrCreate(ctx, *variant)
		if err == nil {
			_ = uc.AuditProvider.Create(ctx, variant.Sku, constants.JobUpdateStock, variant.Stock)
		}
	}

	return nil
}
