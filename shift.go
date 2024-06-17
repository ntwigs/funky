package funky

// Shift removes the first element from an array and returns that removed element. This method changes the length of the array.
func Shift[T any](arr *[]T) T {
	if len(*arr) == 0 {
		var zero T
		return zero
	}
	firstElement := (*arr)[0]
	*arr = (*arr)[1:]
	return firstElement
}
