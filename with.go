package funky

// With returns a new array with the specified value inserted at the specified index.
func With[T any](arr []T, index int, value T) []T {
	result := append([]T(nil), arr...)
	if index >= 0 && index < len(arr) {
		result[index] = value
	}
	return result
}
