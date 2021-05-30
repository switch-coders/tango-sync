package providers

import "context"

type Audit interface {
	Create(ctx context.Context, sku, j string, v interface{}) error
}
