package shared

import "errors"

type AppError struct {
	error
	Message string
}

func NewAppErrorWithMessage(error error, message string) AppError {
	return AppError{error, message}
}

func NewAppError(error error) AppError {
	return AppError{error, ""}
}

// ErrNotFound not found
var ErrNotFound = errors.New("Not found")
