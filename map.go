package funky

// Map creates a new array populated with the results of calling a provided function on every element in the calling array.
func Map[T any, U any](arr []T, mapFunc func(T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = mapFunc(v)
	}
	return result
}
