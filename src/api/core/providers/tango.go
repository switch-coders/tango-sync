package providers

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter/tango"
)

type Tango interface {
	SearchStock(ctx context.Context, f tango.SearchStock) ([]entities.Stock, error)
	SearchPrice(ctx context.Context, f tango.SearchPrice) ([]entities.Price, error)
	Authenticate(ctx context.Context, t string) error
}
