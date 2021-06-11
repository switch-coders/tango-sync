package infrastructure

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/switch-coders/tango-sync/src/api/core/errors/apierrors"
)

// HandlerFunc is the func type for the custom handlers.
type HandlerFunc func(c *gin.Context) *apierrors.APIError

// ErrorWrapper if handlerFunc return a error,then response will be composed from error's information.
func ErrorWrapper(handlerFunc HandlerFunc, c *gin.Context) {
	err := handlerFunc(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

// GetCallerID extracts the caller.id from the "X-Caller-Id" header or the "caller.id" param.
// It converts the caller.id to int, or returns an error if the conversion fails.
func GetCallerID(c *gin.Context) (int64, error) {
	providedID := c.Request.Header.Get("X-Caller-Id")
	if providedID == "" {
		providedID = c.Query("caller.id")
	}

	id, err := strconv.ParseInt(providedID, 10, 64)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetCode(c *gin.Context) (string, error) {
	code := c.Request.Header.Get("code")
	if code == "" {
		return "", fmt.Errorf("fails to get code")
	}

	return code, nil
}

//WithRequestContext returns a context.Context from a gin.Context
func WithRequestContext(ctx context.Context, ginCtx *gin.Context) {
	ginCtx.Set(contextKey.String(), ctx)
}

//RequestContext returns a context.Context from a gin.Context
func RequestContext(ginCtx *gin.Context) context.Context {
	ctxValue, ok := ginCtx.Get(contextKey.String())
	if !ok {
		return context.Background()
	}

	return ctxValue.(context.Context)
}
