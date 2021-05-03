package providers

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter"
)

type Tango interface {
	SearchStock(ctx context.Context, f filter.SearchStock) ([]entities.Stock, error)
}
