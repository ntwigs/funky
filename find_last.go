package funky

// FindLast returns the last element in the array that satisfies the provided testing function. If no values satisfy the testing function, the function returns the zero value for type T.
func FindLast[T any](arr []T, test func(T) bool) T {
	for i := len(arr) - 1; i >= 0; i-- {
		if test(arr[i]) {
			return arr[i]
		}
	}
	var zero T
	return zero
}
