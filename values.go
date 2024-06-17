package funky

// Values returns a slice containing all the values in a map.
func Values[K comparable, V any](obj map[K]V) []V {
	var values []V
	for _, value := range obj {
		values = append(values, value)
	}
	return values
}
