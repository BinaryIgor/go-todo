package user

import (
	"fmt"
	"go-todo/shared"
)

const (
	INVALID_AUTH_TOKEN_ERROR    = "INVALID_AUTH_TOKEN"
	INVALID_USER_NAME_ERROR     = "INVALID_USER_NAME"
	INVALID_USER_PASSWORD_ERROR = "INVALID_USER_PASSWORD"
)

type InvalidAuthTokenError struct {
	shared.AppError
}

func NewInvalidAuthTokenError(message string) InvalidAuthTokenError {
	return InvalidAuthTokenError{shared.NewAppError(INVALID_AUTH_TOKEN_ERROR, message)}
}

func ThrowInvalidAuthTokenError(message string) {
	NewInvalidAuthTokenError(message).Throw()
}

func NewInvalidUserNameError(name string) shared.ValidationError {
	return shared.NewValidationError(INVALID_USER_NAME_ERROR,
		fmt.Sprintln(name, "is not a valid user name.",
			"It has to start with a letter, and can only have alphanumeric characters afterwards.",
			"Min length is 3 and max is 30"))
}

func NewInvalidUserPasswordError() shared.ValidationError {
	return shared.NewValidationError(INVALID_USER_PASSWORD_ERROR,
		"Invalid password. It needs to have at least 8 characters, one digit, one lower and one uppercase letter")
}
