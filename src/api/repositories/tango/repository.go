package tango

import (
	"context"
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/go-resty/resty/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter/tango"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
	"strconv"
)

type Repository struct {
	TangoBaseURL     string
	TangoAccessToken string
}

func (r *Repository) SearchStock(ctx context.Context, f tango.SearchStock) ([]entities.Stock, error) {
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

func (r *Repository) SearchPrice(ctx context.Context, f tango.SearchPrice) ([]entities.Price, error) {
	transaction := newrelic.FromContext(ctx)

	client := resty.New()
	url := r.TangoBaseURL + "/Price"

	var pages []price
	morePages := true
	pageNumber := 1
	pageSize := 250

	var restResp *resty.Response
	var err error
	for morePages {
		infrastructure.WrapExternalSegmentWithAlias(transaction, url, "TANGO_PRICE", func() {
			restResp, err = client.R().
				SetHeader("accesstoken", r.TangoAccessToken).
				SetQueryParams(map[string]string{
					"pageNumber": strconv.Itoa(pageNumber),
					"pageSize":   strconv.Itoa(pageSize),
					"filter":     f.ListPriceNumber,
					"lastUpdate": f.LastUpdate,
				}).Get(url)
		})

		if err != nil {
			sentry.CaptureException(err)

			return nil, errors.NewRepositoryError(errors.ErrorGettingResource.GetMessageWithParams(errors.Parameters{
				"resource": "search-tango-price",
			}))
		}

		if restResp.IsError() {
			return nil, errors.NewDependencyError(restResp.String())
		}

		var resp price
		err = json.Unmarshal(restResp.Body(), &resp)
		if err != nil {
			sentry.CaptureException(err)

			return nil, errors.NewParsingError(errors.ErrorUnmarshallingResponse.GetMessageWithParams(errors.Parameters{
				"resource": "search-tango-price",
			}))
		}

		pages = append(pages, resp)
		if !resp.Paging.MoreData {
			morePages = false
		}

		pageNumber++
	}

	var prices []entities.Price
	for _, p := range pages {
		prices = append(prices, p.GetPriceEntity()...)
	}

	return prices, nil
}

func (r *Repository) Authenticate(ctx context.Context, t string) error {
	transaction := newrelic.FromContext(ctx)

	client := resty.New()
	url := r.TangoBaseURL + "/dummy"

	var restResp *resty.Response
	var err error

	infrastructure.WrapExternalSegmentWithAlias(transaction, url, "TANGO_AUTHENTICATION", func() {
		restResp, err = client.R().
			SetHeader("accesstoken", t).
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json").
			SetBody([]byte{}).
			Post(url)
	})

	if err != nil {
		sentry.CaptureException(err)

		return errors.NewRepositoryError(errors.ErrorGettingResource.GetMessageWithParams(errors.Parameters{
			"resource": "authenticate-tango",
		}))
	}

	if restResp.IsError() {
		return errors.NewDependencyError(restResp.String())
	}

	var resp authentication
	err = json.Unmarshal(restResp.Body(), &resp)
	if err != nil {
		sentry.CaptureException(err)

		return errors.NewParsingError(errors.ErrorUnmarshallingResponse.GetMessageWithParams(errors.Parameters{
			"resource": "authenticate-tango",
		}))
	}

	if !resp.IsOk {
		return errors.NewForbiddenError(errors.ErrorInvalidTangoToken.GetMessageWithParams(errors.Parameters{
			"resource": "authenticate-tango",
		}))
	}

	return nil
}
