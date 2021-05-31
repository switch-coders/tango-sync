package integration

import (
	"context"
	"github.com/switch-coders/tango-sync/src/api/core/contracts/integration"
)

type UseCase interface {
	Execute(ctx context.Context, r integration.Request) error
}