package providers

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
)

type Notification interface {
	Notify(ctx context.Context, n entities.Notification) error
}
