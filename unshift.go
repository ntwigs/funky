package funky

// Unshift adds one or more elements to the beginning of an array and returns the new length of the array.
func Unshift[T any](arr *[]T, elements ...T) int {
	*arr = append(elements, *arr...)
	return len(*arr)
}
