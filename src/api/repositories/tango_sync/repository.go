package tango_sync

import (
	"github.com/getsentry/sentry-go"
	"github.com/go-resty/resty/v2"
)

type Repository struct {
	BaseURL string
}

func (r *Repository) ExecuteStockSync() {
	client := resty.New()

	url := r.BaseURL + "/jobs/sync/stock"

	restResp, err := client.R().
		Post(url)

	if err != nil {
		sentry.CaptureException(err)
	}

	if restResp.IsError() {
		sentry.CaptureMessage(string(restResp.Body()))
	}
}

func (r *Repository) ExecutePriceSync() {
	client := resty.New()

	url := r.BaseURL + "/jobs/sync/price"

	restResp, err := client.R().
		Post(url)

	if err != nil {
		sentry.CaptureException(err)
	}

	if restResp.IsError() {
		sentry.CaptureMessage(string(restResp.Body()))
	}
}
