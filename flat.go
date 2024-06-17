package funky

// Flat returns a new array with all sub-array elements concatenated into it recursively up to the specified depth.
func Flat[T any](arr [][]T) []T {
	var result []T
	for _, v := range arr {
		result = append(result, v...)
	}
	return result
}
