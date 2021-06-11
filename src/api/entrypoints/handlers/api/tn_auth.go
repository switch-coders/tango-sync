package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
	tnAuth "github.com/switch-coders/tango-sync/src/api/core/usecases/tn_auth"
	"github.com/switch-coders/tango-sync/src/api/infrastructure"
)

type TnAuth struct {
	TnAuthUseCase tnAuth.UseCase
}

func (handler *TnAuth) Handle(c *gin.Context) {
	infrastructure.ErrorWrapper(handler.handle, c)
}

const ttl = 21600

func (handler *TnAuth) handle(c *gin.Context) *apierrors.APIError {
	ctx := infrastructure.ContextFrom(c)

	code := c.Query("code")

	account, err := handler.TnAuthUseCase.Execute(ctx, code)
	if err != nil {
		return apierrors.GetCommonsAPIError(err)
	}

	c.SetCookie("tn_token", account.AccessToken, ttl, "", "", true, true)
	c.SetCookie("tn_user_id", strconv.Itoa(int(account.UserID)), ttl, "", "", true, true)
	c.Redirect(http.StatusMovedPermanently, "/register")
	return nil
}
