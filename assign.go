package funky

// Assign assigns properties from source maps to the target map.
func Assign[K comparable, V any](target map[K]V, sources ...map[K]V) map[K]V {
	for _, source := range sources {
		for key, value := range source {
			target[key] = value
		}
	}
	return target
}
