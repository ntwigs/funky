package funky

// CopyWithin shallow copies part of an array to another location in the same array and returns it without modifying its length.
func CopyWithin[T any](arr []T, target, start, end int) []T {
	if start < 0 {
		start = len(arr) + start
	}
	if end < 0 {
		end = len(arr) + end
	}
	if target < 0 {
		target = len(arr) + target
	}

	if start < 0 {
		start = 0
	}
	if end > len(arr) {
		end = len(arr)
	}
	if target >= len(arr) {
		return arr
	}

	copyLen := end - start
	if copyLen <= 0 {
		return arr
	}

	if target+copyLen > len(arr) {
		copyLen = len(arr) - target
	}

	copy(arr[target:target+copyLen], arr[start:start+copyLen])
	return arr
}
