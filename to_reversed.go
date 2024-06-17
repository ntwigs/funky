package funky

// ToReversed returns a new array with the elements in reversed order.
func ToReversed[T any](arr []T) []T {
	result := make([]T, len(arr))
	copy(result, arr)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
