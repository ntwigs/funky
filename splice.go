package funky

// Splice changes the contents of an array by removing or replacing existing elements and/or adding new elements in place.
func Splice[T any](arr *[]T, start, deleteCount int, items ...T) []T {
	if start < 0 {
		start = len(*arr) + start
	}
	if start < 0 {
		start = 0
	} else if start > len(*arr) {
		start = len(*arr)
	}

	if deleteCount < 0 {
		deleteCount = 0
	} else if deleteCount > len(*arr)-start {
		deleteCount = len(*arr) - start
	}

	removed := append([]T(nil), (*arr)[start:start+deleteCount]...)
	*arr = append((*arr)[:start], append(items, (*arr)[start+deleteCount:]...)...)

	return removed
}
