package api

import (
	"github.com/gin-gonic/gin"
	integrationContract "github.com/switch-coders/tango-sync/src/api/core/contracts/integration"
	"github.com/switch-coders/tango-sync/src/api/core/errors"
	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/integration"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
	"net/http"
)

type Registration struct {
	IntegrationUseCase integration.UseCase
}

func (handler *Registration) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(handler.handle, c)
}

func (handler *Registration) handle(c *gin.Context) *apierrors.APIError {
	ctx := infrastructure.ContextFrom(c)

	var request integrationContract.Request
	err := c.ShouldBind(&request)
	if err != nil {
		return apierrors.NewBadRequest(errors.ErrorBindingRequest.GetMessage(), err.Error())
	}

	request.TnUserID, err = c.Cookie("tn_user_id")
	if err != nil {
		return apierrors.NewBadRequest(errors.ErrorBindingRequest.GetMessage(), err.Error())
	}

	request.TnAccessToken, err = c.Cookie("tn_token")
	if err != nil {
		return apierrors.NewBadRequest(errors.ErrorBindingRequest.GetMessage(), err.Error())
	}

	err = handler.IntegrationUseCase.Execute(ctx, request)
	if err != nil {
		return apierrors.GetCommonsAPIError(err)
	}

	c.Status(http.StatusOK)
	return nil
}
