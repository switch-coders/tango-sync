package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	integrationContract "github.com/switch-coders/tango-sync/src/api/core/contracts/integration"
	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/integration"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
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
		c.Redirect(http.StatusMovedPermanently, "/error")
		return nil
	}

	request.TnUserID, err = c.Cookie("tn_user_id")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/error")
		return nil
	}

	request.TnAccessToken, err = c.Cookie("tn_token")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/error")
		return nil
	}

	err = handler.IntegrationUseCase.Execute(ctx, request)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/error")
		return nil
	}

	c.SetCookie("integration_success", "true", 200, "", "", false, false)
	c.Redirect(http.StatusMovedPermanently, "/index")
	return nil
}
