package errors

type DependencyError struct {
	message string
}

func NewDependencyError(message string) DependencyError {
	return DependencyError{
		message: message,
	}
}

func (e DependencyError) Error() string {
	return e.message
}
