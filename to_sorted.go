package funky

import "sort"

// ToSorted returns a new array with the elements sorted in ascending order.
func ToSorted[T any](arr []T, less func(T, T) bool) []T {
	result := make([]T, len(arr))
	copy(result, arr)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}
