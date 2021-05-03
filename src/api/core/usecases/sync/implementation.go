package sync

import (
	"context"
	goErrors "errors"
	"time"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/providers"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Implementation struct {
	TangoProvider        providers.Tango
	ProductProvider      providers.Product
	NotificationProvider providers.Notification
}

func (uc *Implementation) Execute(ctx context.Context, lastUpdate *time.Time) error {
	ctx = context.WithValue(ctx, infrastructure.ActionKey, "sync")

	stocksProduct, err := uc.TangoProvider.SearchStock(ctx, filter.NewSearchStocktango(lastUpdate))
	if err != nil {
		return errors.GetError(err, err.Error())
	}

	for _, stockProduct := range stocksProduct {
		p, err := uc.ProductProvider.Get(ctx, stockProduct.SkuCode)
		if err != nil {
			if goErrors.As(err, &errors.RepositoryError{}) {
				return errors.GetError(err, err.Error())
			}
		} else {
			if stockIsEqual(p.Stock, int(stockProduct.Quantity)) {
				continue
			}
		}

		err = uc.NotificationProvider.Notify(ctx, entities.NewSkuStockNotification(stockProduct.SkuCode, stockProduct.Quantity))
		if err != nil {
			return errors.GetError(err, err.Error())
		}
	}

	return nil
}

func stockIsEqual(q1, q2 int) bool {
	return q1 == q2
}
