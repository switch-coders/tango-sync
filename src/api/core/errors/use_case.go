package errors

import "strings"

func GetError(err error, messages ...string) error {
	switch err.(type) {
	case BadRequestError:
		return NewBadRequestError(strings.Join(messages, " "))
	case RepositoryError, InternalServerError:
		return NewInternalServerError(strings.Join(messages, " "))
	case DependencyError:
		return NewDependencyError(strings.Join(messages, " "))
	case ForbiddenError:
		return NewForbiddenError(strings.Join(messages, " "))
	}
	return err
}
