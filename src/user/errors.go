package user

import (
	"go-todo/shared"
)

const INVALID_AUTH_TOKEN_ERROR_CODE = "INVALID_AUTH_TOKEN_ERROR"

type InvalidAuthTokenError struct {
	shared.AppError
}

func NewInvalidAuthTokenError(message string) InvalidAuthTokenError {
	return InvalidAuthTokenError{shared.AppError{INVALID_AUTH_TOKEN_ERROR_CODE, message}}
}

func ThrowInvalidAuthTokenError(message string) {
	NewInvalidAuthTokenError(message).Throw()
}
