package shared

import "strings"

type ApiError struct {
	Errors  []string `json:"errors"`
	Message string   `json:"message"`
}

func NewApiErrorWithMultipleErrors(errors []string, message string) ApiError {
	return ApiError{errors, message}
}

func NewApiError(error string, message string) ApiError {
	return ApiError{[]string{error}, message}
}

func IsEndpointPublic(path string) bool {
	return strings.HasPrefix(path, "/users/sign-up") ||
		strings.HasPrefix(path, "/users/sign-in")
}
