package update_price

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/contracts/sync_by_product/price"
)

type UseCase interface {
	Execute(ctx context.Context, request price.Request) error
}
