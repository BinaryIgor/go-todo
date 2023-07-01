package shared

func MapSlice[T, U any](collection []T, mapping func(T) U) []U {
	mapped := make([]U, len(collection))
	for i := range collection {
		mapped[i] = mapping(collection[i])
	}
	return mapped
}
