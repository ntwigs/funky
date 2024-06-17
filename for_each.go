package funky

// ForEach executes a provided function once for each array element.
func ForEach[T any](arr []T, callback func(T)) {
	for _, v := range arr {
		callback(v)
	}
}
