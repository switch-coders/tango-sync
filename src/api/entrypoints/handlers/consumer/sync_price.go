package consumer

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"

	"github.com/switch-coders/tango-sync/src/api/core/contracts/sync_by_product/price"
	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/update_price"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type UpdatePrice struct {
	UpdatePriceUseCase update_price.UseCase
}

func (handler *UpdatePrice) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(handler.handle, c)
}

func (handler *UpdatePrice) handle(c *gin.Context) *apierrors.APIError {
	ctx := infrastructure.ContextFrom(c)

	var request price.Request
	err := c.BindJSON(&request)
	if err != nil {
		sentry.CaptureException(err)
		return apierrors.GetCommonsAPIError(err)
	}

	err = handler.UpdatePriceUseCase.Execute(ctx, request)
	if err != nil {
		return apierrors.GetCommonsAPIError(err)
	}
	c.Status(http.StatusOK)

	return nil
}
