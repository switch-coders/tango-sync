package update_stock

import (
	"context"
	"github.com/switch-coders/tango-sync/src/api/core/contracts/sync_by_product/stock"
)

type UseCase interface {
	Execute(ctx context.Context, request stock.Request) error
}
