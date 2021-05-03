package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/core/usecases/get"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Get struct {
	GetUseCase get.UseCase
}

func (handler *Get) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(handler.handle, c)
}

func (handler *Get) handle(c *gin.Context) *apierrors.APIError {

	value, _ := handler.GetUseCase.Execute(c)
	c.JSON(http.StatusOK, value)

	return nil
}
