package funky

// ToSpliced returns a new array with the specified elements spliced.
func ToSpliced[T any](arr []T, start, deleteCount int, items ...T) []T {
	if start < 0 {
		start = len(arr) + start
	}
	if start < 0 {
		start = 0
	} else if start > len(arr) {
		start = len(arr)
	}

	if deleteCount < 0 {
		deleteCount = 0
	} else if deleteCount > len(arr)-start {
		deleteCount = len(arr) - start
	}

	result := append([]T(nil), arr[:start]...)
	result = append(result, items...)
	result = append(result, arr[start+deleteCount:]...)
	return result
}
