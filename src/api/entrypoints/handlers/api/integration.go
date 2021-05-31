package api

import (
	"github.com/gin-gonic/gin"
	integrationContract "github.com/switch-coders/tango-sync/src/api/core/contracts/integration"
	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/integration"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
	"net/http"
)

type Integration struct {
	IntegrationUseCase integration.UseCase
}

func (handler *Integration) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(handler.handle, c)
}

func (handler *Integration) handle(c *gin.Context) *apierrors.APIError {
	ctx := infrastructure.ContextFrom(c)

	var request integrationContract.Request
	err := c.Bind(&request)
	if err != nil {
		return apierrors.GetCommonsAPIError(err)
	}

	err = handler.IntegrationUseCase.Execute(ctx, request)
	if err != nil {
		return apierrors.GetCommonsAPIError(err)
	}

	c.Header("tn_user_id", "153223")
	c.Header("tn_access_token", "23f1bbf0b7c70a0f494bc79f08a435633267c940")

	c.Redirect(http.StatusMovedPermanently, "https://www.tiendanube.com/apps/2736/authorize")

	return nil
}
