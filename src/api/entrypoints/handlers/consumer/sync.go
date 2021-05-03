package consumer

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	syncByProduct "github.com/switch-coders/tango-sync/src/api/core/contracts/sync_by_product"
	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/sync_by_product"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type SyncByProduct struct {
	SyncByProductUseCase sync_by_product.UseCase
}

func (handler *SyncByProduct) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(handler.handle, c)
}

func (handler *SyncByProduct) handle(c *gin.Context) *apierrors.APIError {
	ctx := infrastructure.ContextFrom(c)

	var request syncByProduct.Request
	err := c.BindJSON(&request)
	if err != nil {
		sentry.CaptureException(err)
		return apierrors.GetCommonsAPIError(err)
	}

	err = handler.SyncByProductUseCase.Execute(ctx, request)
	if err != nil {
		return apierrors.GetCommonsAPIError(err)
	}
	c.Status(http.StatusOK)

	return nil
}
