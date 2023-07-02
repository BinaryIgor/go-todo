package user

import (
	"fmt"
	"go-todo/shared"
)

const (
	INVALID_USER_NAME_ERROR     = "INVALID_USER_NAME"
	INVALID_USER_PASSWORD_ERROR = "INVALID_USER_PASSWORD"
	USER_NOT_FOUND_ERRROR       = "USER_NOT_FOUND"
)

var InvalidUserPasswordError = shared.NewValidationError(INVALID_USER_PASSWORD_ERROR,
	"Invalid password. It needs to have at least 8 characters, one digit, one lower and one uppercase letter")

func NewInvalidUserNameError(name string) shared.ValidationError {
	return shared.NewValidationError(INVALID_USER_NAME_ERROR,
		fmt.Sprintln(name, "is not a valid user name.",
			"It has to start with a letter, and can only have alphanumeric characters afterwards.",
			"Min length is 3 and max is 30"))
}

func NewUserNotFoundError(name string) shared.NotFoundError {
	return shared.NewNotFoundError(USER_NOT_FOUND_ERRROR,
		fmt.Sprintf("User of %s name doesn't exist", name))
}
