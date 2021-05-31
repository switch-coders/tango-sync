package api

import (
	"github.com/gin-gonic/gin"
	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
	"net/http"
)

type Integration struct{}

func (handler *Integration) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(handler.handle, c)
}

func (handler *Integration) handle(c *gin.Context) *apierrors.APIError {
	c.Header("tn_user_id", "153223")
	c.Header("tn_access_token", "23f1bbf0b7c70a0f494bc79f08a435633267c940")

	c.Redirect(http.StatusMovedPermanently, "https://www.tiendanube.com/apps/2736/authorize")

	return nil
}
