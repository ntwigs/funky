package funky

import (
	"fmt"
	"strings"
)

// Join creates and returns a new string by concatenating all of the elements in an array (or an array-like object), separated by commas or a specified separator string.
func Join[T any](arr []T, separator string) string {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		strArr[i] = fmt.Sprint(v)
	}
	return strings.Join(strArr, separator)
}
