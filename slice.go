package funky

// Slice returns a shallow copy of a portion of an array into a new array object selected from start to end (end not included) where start and end represent the index of items in that array.
func Slice[T any](arr []T, start, end int) []T {
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
		return []T{}
	}
	return arr[start:end]
}
