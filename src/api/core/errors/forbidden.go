package errors

type ForbiddenError struct {
	message string
}

func NewForbiddenError(message string) ForbiddenError {
	return ForbiddenError{
		message: message,
	}
}

func (e ForbiddenError) Error() string {
	return e.message
}
