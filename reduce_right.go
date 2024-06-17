package funky

// ReduceRight applies a function against an accumulator and each value of the array (from right-to-left) to reduce it to a single value.
func ReduceRight[T any, U any](arr []T, reducer func(U, T) U, initialValue U) U {
	result := initialValue
	for i := len(arr) - 1; i >= 0; i-- {
		result = reducer(result, arr[i])
	}
	return result
}
