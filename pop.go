package funky

// Pop removes the last element from an array and returns that element. This method changes the length of the array.
func Pop[T any](arr *[]T) T {
	if len(*arr) == 0 {
		var zero T
		return zero
	}
	lastElement := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return lastElement
}
