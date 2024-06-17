package funky

// Fill changes all elements in an array to a static value, from a start index (default 0) to an end index (default array length). It returns the modified array.
func Fill[T any](arr []T, value T, start, end int) []T {
	if start < 0 {
		start = len(arr) + start
	}
	if end < 0 {
		end = len(arr) + end
	}

	if start < 0 {
		start = 0
	}
	if end > len(arr) {
		end = len(arr)
	}
	if start > end {
		start = end
	}

	for i := start; i < end; i++ {
		arr[i] = value
	}
	return arr
}
