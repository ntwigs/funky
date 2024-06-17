package funky

// Push adds one or more elements to the end of an array and returns the new length of the array.
func Push[T any](arr *[]T, elements ...T) int {
	*arr = append(*arr, elements...)
	return len(*arr)
}
