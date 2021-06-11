package providers

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/contracts/integration"
)

type Account interface {
	Create(ctx context.Context, r integration.Request) error
}
