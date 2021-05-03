package sync_by_product

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/contracts/sync_by_product"
)

type UseCase interface {
	Execute(ctx context.Context, request sync_by_product.Request) error
}
