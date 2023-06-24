package shared

import "strings"

type ApiError struct {
	errors  []string
	message string
}

func NewApiError(errors []string, message string) ApiError {
	return ApiError{errors, message}
}

func IsEndpointPublic(path string) bool {
	return strings.HasPrefix(path, "/users/sign-up") ||
		strings.HasPrefix(path, "/users/sign-in")
}
