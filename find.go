package funky

// Find returns the first element in the array that satisfies the provided testing function. If no values satisfy the testing function, the function returns the zero value for type T.
func Find[T any](arr []T, test func(T) bool) T {
	for _, v := range arr {
		if test(v) {
			return v
		}
	}
	var zero T
	return zero
}
