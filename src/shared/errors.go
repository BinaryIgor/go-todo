package shared

type AppError interface {
	error
	Code() string
	Message() string
	Throw()
}

type baseAppError struct {
	code    string
	message string
}

func (e *baseAppError) Error() string {
	return e.message
}

func (e *baseAppError) Code() string {
	return e.code
}

func (e *baseAppError) Message() string {
	return e.message
}

func (e *baseAppError) Throw() {
	panic(e)
}

func NewAppError(code string, message string) AppError {
	return &baseAppError{code, message}
}

type ValidationError struct {
	AppError
}

func NewValidationError(code string, message string) ValidationError {
	return ValidationError{&baseAppError{code, message}}
}

func ThrowValidationError(code string, message string) {
	NewValidationError(code, message).Throw()
}

type NotFoundError struct {
	AppError
}

func NewNotFoundError(code string, message string) NotFoundError {
	return NotFoundError{&baseAppError{code, message}}
}
