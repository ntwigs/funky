package funky

// IndexOf returns the first index at which a given element can be found in the array, or -1 if it is not present.
func IndexOf[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}
