package tn_auth

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/providers"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Implementation struct {
	TnProvider providers.TiendaNube
}

func (uc *Implementation) Execute(ctx context.Context, code string) (*entities.TnAccount, error) {
	ctx = context.WithValue(ctx, infrastructure.ActionKey, "tn_auth")

	account, err := uc.TnProvider.Authorize(ctx, code)
	if err != nil {
		return nil, errors.GetError(err, err.Error())
	}

	return account, nil
}
