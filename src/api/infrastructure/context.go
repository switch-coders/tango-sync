package infrastructure

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type ctxKey string

func (c ctxKey) String() string {
	return "sync_context_key_" + string(c)
}

const (
	ActionKey  = ctxKey("action")
	contextKey = ctxKey("context_key")
)

func ContextFrom(c *gin.Context) context.Context {
	ctxValue, ok := c.Get(contextKey.String())
	if !ok {
		return context.Background()
	}

	return ctxValue.(context.Context)
}

func SetNRTransaction(c *gin.Context) {
	txn := nrgin.Transaction(c)
	if txn != nil {
		name := transactionName(c)
		txn.SetName(name)
	}
	requestCtx := RequestContext(c)
	newCtx := newrelic.NewContext(requestCtx, txn)
	WithRequestContext(newCtx, c)
	c.Next()
}

func transactionName(c *gin.Context) string {
	name := c.Request.URL.Path

	stringArray := strings.FieldsFunc(name, func(r rune) bool { return r == '/' })

	_, err := strconv.ParseInt(stringArray[len(stringArray)-1], 10, 64)

	if err == nil {
		stringArray = stringArray[:len(stringArray)-1]
		name = fmt.Sprintf("%s - /%s", c.Request.Method, strings.Join(stringArray, "/"))
	}

	return name
}
