package funky

// Includes determines whether an array includes a certain value among its entries, returning true or false as appropriate.
func Includes[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
