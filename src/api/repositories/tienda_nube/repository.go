package tienda_nube

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/go-resty/resty/v2"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/switch-coders/tango-sync/src/api/core/entities"
	"github.com/switch-coders/tango-sync/src/api/core/entities/filter"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Repository struct {
	TNBaseURL        string
	TNAuthentication string
	TNUserAgent      string
	TNNumber         string
	TNSecret         string
	TNAppID          string
}

func (r *Repository) Authorize(ctx context.Context, code string) (*entities.TnAccount, error) {
	transaction := newrelic.FromContext(ctx)

	client := resty.New()

	url := fmt.Sprintf("https://www.tiendanube.com/apps/authorize/token")

	var restResp *resty.Response
	var err error

	infrastructure.WrapExternalSegmentWithAlias(transaction, url, "AUTHORIZE_TN_ACCOUNT", func() {
		restResp, err = client.R().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetFormData(map[string]string{
				"client_id":     r.TNAppID,
				"client_secret": r.TNSecret,
				"grant_type":    "authorization_code",
				"code":          code,
			}).
			Post(url)
	})

	if err != nil {
		sentry.CaptureException(err)

		return nil, errors.NewRepositoryError(errors.ErrorGettingResource.GetMessageWithParams(errors.Parameters{
			"resource": "search-tn-product",
		}))
	}

	if restResp.IsError() {
		return nil, errors.NewRepositoryError(restResp.String())
	}

	var resp account
	err = json.Unmarshal(restResp.Body(), &resp)
	if err != nil {
		sentry.CaptureException(err)

		return nil, errors.NewParsingError(errors.ErrorUnmarshallingResponse.GetMessageWithParams(errors.Parameters{
			"resource": "search-tn-product",
		}))
	}

	return resp.toEntity(), nil
}

func (r *Repository) SearchProduct(ctx context.Context, filter filter.SearchProduct) (*entities.Product, error) {
	transaction := newrelic.FromContext(ctx)

	client := resty.New()

	url := fmt.Sprintf(r.TNBaseURL+"/products", r.TNNumber)

	var restResp *resty.Response
	var err error

	infrastructure.WrapExternalSegmentWithAlias(transaction, url, "SEARCH_TN_PRODUCT", func() {
		restResp, err = client.R().
			SetHeader("Authentication", r.TNAuthentication).
			SetHeader("User-Agent", r.TNUserAgent).
			SetQueryParams(map[string]string{
				"q":      filter.Q,
				"fields": strings.Join(filter.Fields, ","),
			}).
			Get(url)
	})

	if err != nil {
		sentry.CaptureException(err)

		return nil, errors.NewRepositoryError(errors.ErrorGettingResource.GetMessageWithParams(errors.Parameters{
			"resource": "search-tn-product",
		}))
	}

	if restResp.IsError() {
		if restResp.StatusCode() == http.StatusTooManyRequests {
			sentry.CaptureMessage(errors.ErrorTooManyRequests.GetMessage())
		}

		if restResp.StatusCode() == http.StatusNotFound {
			return nil, errors.NewNotFoundError(restResp.String())
		}

		return nil, errors.NewRepositoryError(restResp.String())
	}

	var resp products
	err = json.Unmarshal(restResp.Body(), &resp)
	if err != nil {
		sentry.CaptureException(err)

		return nil, errors.NewParsingError(errors.ErrorUnmarshallingResponse.GetMessageWithParams(errors.Parameters{
			"resource": "search-tn-product",
		}))
	}

	return resp.toEntity(filter.Q), nil
}

func (r *Repository) UpdateProductStockVariant(ctx context.Context, product entities.Product) error {
	transaction := newrelic.FromContext(ctx)

	client := resty.New()

	url := fmt.Sprintf(r.TNBaseURL+"/products/%d/variants/%d", r.TNNumber, product.ProductID, product.ID)

	var restResp *resty.Response
	var err error

	infrastructure.WrapExternalSegmentWithAlias(transaction, url, "PUT_TN_PRODUCT", func() {
		restResp, err = client.R().
			SetHeader("Authentication", r.TNAuthentication).
			SetHeader("User-Agent", r.TNUserAgent).
			SetBody(map[string]interface{}{"stock": product.Stock}).
			Put(url)
	})

	if err != nil {
		sentry.CaptureException(err)

		return errors.NewRepositoryError(errors.ErrorUpdatingResource.GetMessageWithParams(errors.Parameters{
			"resource": "product",
			"action":   "update",
			"sku":      product.Sku,
		}))
	}

	if restResp.IsError() {
		return errors.NewRepositoryError(restResp.String())
	}

	var resp variants
	err = json.Unmarshal(restResp.Body(), &resp)
	if err != nil {
		sentry.CaptureException(err)

		return errors.NewParsingError(errors.ErrorUnmarshallingResponse.GetMessageWithParams(errors.Parameters{
			"resource": "product",
			"action":   "update",
			"sku":      product.Sku,
		}))
	}

	return nil
}

func (r *Repository) UpdateProductPriceVariant(ctx context.Context, product entities.Product) error {
	transaction := newrelic.FromContext(ctx)

	client := resty.New()

	url := fmt.Sprintf(r.TNBaseURL+"/products/%d/variants/%d", r.TNNumber, product.ProductID, product.ID)

	var restResp *resty.Response
	var err error

	infrastructure.WrapExternalSegmentWithAlias(transaction, url, "PUT_TN_PRODUCT", func() {
		restResp, err = client.R().
			SetHeader("Authentication", r.TNAuthentication).
			SetHeader("User-Agent", r.TNUserAgent).
			SetBody(map[string]interface{}{"price": product.Price}).
			Put(url)
	})

	if err != nil {
		sentry.CaptureException(err)

		return errors.NewRepositoryError(errors.ErrorUpdatingResource.GetMessageWithParams(errors.Parameters{
			"resource": "product",
			"action":   "update",
			"sku":      product.Sku,
		}))
	}

	if restResp.IsError() {
		return errors.NewRepositoryError(restResp.String())
	}

	var resp variants
	err = json.Unmarshal(restResp.Body(), &resp)
	if err != nil {
		sentry.CaptureException(err)

		return errors.NewParsingError(errors.ErrorUnmarshallingResponse.GetMessageWithParams(errors.Parameters{
			"resource": "product",
			"action":   "update",
			"sku":      product.Sku,
		}))
	}

	return nil
}
