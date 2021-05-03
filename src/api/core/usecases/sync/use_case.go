package sync

import (
	"context"
	"time"
)

type UseCase interface {
	Execute(ctx context.Context, lastUpdate *time.Time) error
}
