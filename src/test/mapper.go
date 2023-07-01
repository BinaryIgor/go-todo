package test

import "go-todo/shared"

func ApiErrorFromAppError(error shared.AppError) shared.ApiError {
	return shared.NewApiError(error.Code(), error.Message())
}
