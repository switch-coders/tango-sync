package providers

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
)

type Product interface {
	UpdateOrCreate(ctx context.Context, product entities.Product) error
	Get(ctx context.Context, sku string) (*entities.Product, error)
}
