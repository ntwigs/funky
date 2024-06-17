package funky

// Every tests whether all elements in the array pass the test implemented by the provided function.
func Every[T any](arr []T, test func(T) bool) bool {
	for _, v := range arr {
		if !test(v) {
			return false
		}
	}
	return true
}
