package funky

// Filter creates a new array with all elements that pass the test implemented by the provided function.
func Filter[T any](arr []T, test func(T) bool) []T {
	var result []T
	for _, v := range arr {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}
