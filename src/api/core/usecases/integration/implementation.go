package integration

import (
	"context"

	"github.com/switch-coders/tango-sync/src/api/core/contracts/integration"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/providers"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Implementation struct {
	TangoProvider   providers.Tango
	AccountProvider providers.Account
}

func (uc *Implementation) Execute(ctx context.Context, r integration.Request) error {
	ctx = context.WithValue(ctx, infrastructure.ActionKey, "tango_integration")

	err := uc.TangoProvider.Authenticate(ctx, r.TangoKey)
	if err != nil {
		return errors.GetError(err, err.Error())
	}

	err = uc.AccountProvider.Create(ctx, r)
	if err != nil {
		return errors.GetError(err, err.Error())
	}

	return nil
}
