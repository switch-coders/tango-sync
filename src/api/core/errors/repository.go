package errors

type RepositoryError struct {
	message string
}

func NewRepositoryError(message string) RepositoryError {
	return RepositoryError{
		message: message,
	}
}

func (e RepositoryError) Error() string {
	return e.message
}
