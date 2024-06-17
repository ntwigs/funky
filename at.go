package funky

// At returns the element at the specified index.
func At[T any](arr []T, index int) T {
	if index < 0 {
		index = len(arr) + index
	}
	if index < 0 || index >= len(arr) {
		var zero T
		return zero
	}
	return arr[index]
}
