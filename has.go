package funky

// Has checks if a map contains a specific key. Returns true if the key exists in the map, false otherwise.
func Has[K comparable, V any](obj map[K]V, key K) bool {
	_, exists := obj[key]
	return exists
}
