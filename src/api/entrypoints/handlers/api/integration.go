package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type Integration struct {
	TnAppID string
}

func (h *Integration) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(h.handle, c)
}

const path = "https://www.tiendanube.com/apps/%s/authorize"

func (h *Integration) handle(c *gin.Context) *apierrors.APIError {
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf(path, h.TnAppID))

	return nil
}
