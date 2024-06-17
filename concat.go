package funky

// Concat returns a new array that is the result of concatenating the provided arrays.
func Concat[T any](arr []T, values ...[]T) []T {
	result := make([]T, len(arr))
	copy(result, arr)
	for _, v := range values {
		result = append(result, v...)
	}
	return result
}
