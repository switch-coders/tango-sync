package jobs

import (
	"github.com/gin-gonic/gin"
	syncContract "github.com/switch-coders/tango-sync/src/api/core/contracts/sync"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/sync"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
	"net/http"
)

type SyncStock struct {
	SyncUseCase sync.UseCase
}

func (handler *SyncStock) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(handler.handle, c)
}

func (handler *SyncStock) handle(c *gin.Context) *apierrors.APIError {
	ctx := infrastructure.ContextFrom(c)

	var r syncContract.Request
	err := c.ShouldBind(&r)
	if err != nil {
		return apierrors.NewBadRequest(errors.ErrorBindingRequest.GetMessage(), err.Error())
	}

	err = handler.SyncUseCase.Stock(ctx, r.LastUpdate)
	if err != nil {
		return apierrors.GetCommonsAPIError(err)
	}
	c.Status(http.StatusOK)

	return nil
}
