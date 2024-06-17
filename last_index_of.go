package funky

// LastIndexOf returns the last index at which a given element can be found in the array, or -1 if it is not present.
func LastIndexOf[T comparable](arr []T, value T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == value {
			return i
		}
	}
	return -1
}
