package tn_auth

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
)

type UseCase interface {
	Execute(ctx context.Context, code string) (*entities.TnAccount, error)
}
