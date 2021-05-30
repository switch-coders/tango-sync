package sync

import (
	"context"
	"time"
)

type UseCase interface {
	Stock(ctx context.Context, lastUpdate *time.Time) error
	Price(ctx context.Context, lastUpdate *time.Time) error
}
