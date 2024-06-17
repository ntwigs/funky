package funky

// Reduce executes a reducer function (that you provide) on each element of the array, resulting in a single output value.
func Reduce[T any, U any](arr []T, reducer func(U, T) U, initialValue U) U {
	result := initialValue
	for _, v := range arr {
		result = reducer(result, v)
	}
	return result
}
