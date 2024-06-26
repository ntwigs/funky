package funky

// Keys returns a slice containing all the keys in a map.
// The keys are not always in order.
// https://go.dev/blog/maps
func Keys[K comparable, V any](obj map[K]V) []K {
	var keys []K
	for key := range obj {
		keys = append(keys, key)
	}
	return keys
}
