package providers

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter"
)

type TiendaNube interface {
	UpdateProductStockVariant(ctx context.Context, product entities.Product) error
	UpdateProductPriceVariant(ctx context.Context, product entities.Product) error
	SearchProduct(ctx context.Context, filter filter.SearchProduct) (*entities.Product, error)
}