package apierrors

import (
	"net/http"

	"github.com/switch-coders/tango-sync/src/api/core/errors"
)

func GetCommonsAPIError(err error) *APIError {
	switch err.(type) {
	case errors.BadRequestError:
		return &APIError{Status: http.StatusBadRequest, Message: err.Error(), Err: "bad_request_error", Cause: make([]string, 0)}
	case errors.DependencyError:
		return &APIError{Status: http.StatusFailedDependency, Message: err.Error(), Err: "dependency_error", Cause: make([]string, 0)}
	default:
		return NewInternalServerError(err.Error())
	}
}
