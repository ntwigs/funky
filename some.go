package funky

// Some tests whether at least one element in the array passes the test implemented by the provided function.
func Some[T any](arr []T, test func(T) bool) bool {
	for _, v := range arr {
		if test(v) {
			return true
		}
	}
	return false
}
