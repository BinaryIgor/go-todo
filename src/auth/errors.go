package auth

import (
	"go-todo/shared"
)

const (
	INVALID_AUTH_TOKEN_ERROR = "INVALID_AUTH_TOKEN"
	EXPIRED_AUTH_TOKEN_ERROR = "EXPIRED_AUTH_TOKEN"
)

var InvalidAuthTokenError = shared.NewValidationError(INVALID_AUTH_TOKEN_ERROR, "Auth token is not correct")
var ExpiredAuthTokenError = shared.NewValidationError(EXPIRED_AUTH_TOKEN_ERROR, "Token is expired")
