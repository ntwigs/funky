package funky

// FindIndex returns the index of the first element in the array that satisfies the provided testing function. Otherwise, it returns -1, indicating that no element passed the test.
func FindIndex[T any](arr []T, test func(T) bool) int {
	for i, v := range arr {
		if test(v) {
			return i
		}
	}
	return -1
}
