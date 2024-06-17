package funky

// FlatMap first maps each element using a mapping function, then flattens the result into a new array.
func FlatMap[T any, U any](arr []T, mapFunc func(T) []U) []U {
	var result []U
	for _, v := range arr {
		result = append(result, mapFunc(v)...)
	}
	return result
}
