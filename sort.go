package funky

import "sort"

// Sort sorts the elements of an array in place and returns the array. The default sort order is ascending.
func Sort[T any](arr []T, less func(T, T) bool) []T {
	sort.Slice(arr, func(i, j int) bool {
		return less(arr[i], arr[j])
	})
	return arr
}
