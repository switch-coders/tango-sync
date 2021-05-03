package errors

type ParsingError struct {
	message string
}

func NewParsingError(message string) ParsingError {
	return ParsingError{
		message: message,
	}
}

func (e ParsingError) Error() string {
	return e.message
}
