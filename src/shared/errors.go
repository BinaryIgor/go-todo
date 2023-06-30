package shared

type AppError struct {
	Code    string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) Throw() {
	panic(e)
}

type ValidationError struct {
	AppError
}

func NewValidationError(code string, message string) ValidationError {
	return ValidationError{AppError{code, message}}
}

func ThrowValidationError(code string, message string) {
	NewValidationError(code, message).Throw()
}

type NotFoundError struct {
	AppError
}

func NewNotFoundError(code string, message string) NotFoundError {
	return NotFoundError{AppError{code, message}}
}
