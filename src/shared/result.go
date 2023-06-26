package shared

type Result[T any] struct {
	value *T
	error AppError
}

type EmptyResult struct {
	Result[bool]
}

type emptyPlaceholder struct {
	placeholder bool
}

var empty = emptyPlaceholder{true}

func SuccessEmptyResult() EmptyResult {
	var placeholder *bool
	result := EmptyResult{}
	result.value = placeholder
	return result
}

func SuccessResult[T any](value T) Result[T] {
	return Result[T]{value: &value}
}

func FailureResult[T any](error AppError) Result[T] {
	return Result[T]{nil, error}
}

func (r *Result[T]) IsSuccess() bool {
	return r.value != nil
}

func (r *Result[T]) IsFailure() bool {
	return r.value == nil
}

func (r *Result[T]) Value() T {
	return *r.value
}

func (r *Result[T]) Error() AppError {
	return r.error
}
