package funky

// FindLastIndex returns the index of the last element in the array that satisfies the provided testing function. Otherwise, it returns -1, indicating that no element passed the test.
func FindLastIndex[T any](arr []T, test func(T) bool) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if test(arr[i]) {
			return i
		}
	}
	return -1
}
