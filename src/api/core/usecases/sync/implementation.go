package sync

import (
	"context"
	goErrors "errors"
	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter/tango"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/providers"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
	"github.com/switch-coders/tango-sync/src/api/util/price"
	"github.com/switch-coders/tango-sync/src/api/util/stock"
	"time"
)

type Implementation struct {
	TangoProvider        providers.Tango
	ProductProvider      providers.Product
	NotificationProvider providers.Notification
}

func (uc *Implementation) Stock(ctx context.Context, lastUpdate *time.Time) error {
	ctx = context.WithValue(ctx, infrastructure.ActionKey, "sync_stock")

	stocksProduct, err := uc.TangoProvider.SearchStock(ctx, tango.NewSearchStockEMC(lastUpdate))
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
			if stock.IsEqual(p.Stock, int(stockProduct.Quantity)) {
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

func (uc *Implementation) Price(ctx context.Context, lastUpdate *time.Time) error {
	ctx = context.WithValue(ctx, infrastructure.ActionKey, "sync_price")

	pricesProduct, err := uc.TangoProvider.SearchPrice(ctx, tango.NewSearchPriceEMC(lastUpdate))
	if err != nil {
		return errors.GetError(err, err.Error())
	}

	for _, priceProduct := range pricesProduct {
		p, err := uc.ProductProvider.Get(ctx, priceProduct.SkuCode)
		if err != nil {
			if goErrors.As(err, &errors.RepositoryError{}) {
				return errors.GetError(err, err.Error())
			}
		} else {
			if price.IsEqual(p.Price, priceProduct.Price) {
				continue
			}
		}

		err = uc.NotificationProvider.Notify(ctx, entities.NewSkuPriceNotification(priceProduct.SkuCode, priceProduct.Price))
		if err != nil {
			return errors.GetError(err, err.Error())
		}
	}

	return nil
}
