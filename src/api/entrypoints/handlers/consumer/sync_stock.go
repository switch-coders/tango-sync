package consumer

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"

	"github.com/switch-coders/tango-sync/src/api/core/contracts/sync_by_product/stock"
	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/update_stock"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type UpdateStock struct {
	UpdateStockUseCase update_stock.UseCase
}

func (handler *UpdateStock) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(handler.handle, c)
}

func (handler *UpdateStock) handle(c *gin.Context) *apierrors.APIError {
	ctx := infrastructure.ContextFrom(c)

	var request stock.Request
	err := c.BindJSON(&request)
	if err != nil {
		sentry.CaptureException(err)
		return apierrors.GetCommonsAPIError(err)
	}

	err = handler.UpdateStockUseCase.Execute(ctx, request)
	if err != nil {
		return apierrors.GetCommonsAPIError(err)
	}
	c.Status(http.StatusOK)

	return nil
}
