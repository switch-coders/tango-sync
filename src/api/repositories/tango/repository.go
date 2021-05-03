package tango

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/go-resty/resty/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Repository struct {
	TangoBaseURL     string
	TangoAccessToken string
}

func (r *Repository) SearchStock(ctx context.Context, f filter.SearchStock) ([]entities.Stock, error) {
	transaction := newrelic.FromContext(ctx)

	client := resty.New()
	url := r.TangoBaseURL + "/Stock"

	var pages []stock
	morePages := true
	pageNumber := 1
	pageSize := 250

	var restResp *resty.Response
	var err error
	for morePages {
		infrastructure.WrapExternalSegmentWithAlias(transaction, url, "TANGO_STOCK", func() {
			restResp, err = client.R().
				SetHeader("accesstoken", r.TangoAccessToken).
				SetQueryParams(map[string]string{
					"pageNumber":            strconv.Itoa(pageNumber),
					"pageSize":              strconv.Itoa(pageSize),
					"warehouseCode":         f.WareHouseCode,
					"lastUpdate":            f.LastUpdate,
					"discountPendingOrders": f.DiscountPendingOrders,
				}).Get(url)
		})

		if err != nil {
			sentry.CaptureException(err)

			return nil, errors.NewRepositoryError(errors.ErrorGettingResource.GetMessageWithParams(errors.Parameters{
				"resource": "search-tango-stock",
			}))
		}

		if restResp.IsError() {
			sentry.CaptureException(err)

			return nil, errors.NewDependencyError(restResp.String())
		}

		var resp stock
		err = json.Unmarshal(restResp.Body(), &resp)
		if err != nil {
			sentry.CaptureException(err)

			return nil, errors.NewParsingError(errors.ErrorUnmarshallingResponse.GetMessageWithParams(errors.Parameters{
				"resource": "search-tango-stock",
			}))
		}

		pages = append(pages, resp)
		if !resp.Paging.MoreData {
			morePages = false
		}

		pageNumber++
	}

	var products []entities.Stock
	for _, p := range pages {
		products = append(products, p.GetEntity()...)
	}

	return products, nil
}
