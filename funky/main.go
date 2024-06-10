package funky

import (
	"fmt"
	"sort"
	"strings"
)

// At returns the element at the specified index.
func At[T any](arr []T, index int) T {
	if index < 0 {
		index = len(arr) + index
	}
	if index < 0 || index >= len(arr) {
		var zero T
		return zero
	}
	return arr[index]
}

// Concat returns a new array that is the result of concatenating the provided arrays.
func Concat[T any](arr []T, values ...[]T) []T {
	result := make([]T, len(arr))
	copy(result, arr)
	for _, v := range values {
		result = append(result, v...)
	}
	return result
}

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

// Every tests whether all elements in the array pass the test implemented by the provided function.
func Every[T any](arr []T, test func(T) bool) bool {
	for _, v := range arr {
		if !test(v) {
			return false
		}
	}
	return true
}

// Fill changes all elements in an array to a static value, from a start index (default 0) to an end index (default array length). It returns the modified array.
func Fill[T any](arr []T, value T, start, end int) []T {
	if start < 0 {
		start = len(arr) + start
	}
	if end < 0 {
		end = len(arr) + end
	}

	if start < 0 {
		start = 0
	}
	if end > len(arr) {
		end = len(arr)
	}
	if start > end {
		start = end
	}

	for i := start; i < end; i++ {
		arr[i] = value
	}
	return arr
}

// Filter creates a new array with all elements that pass the test implemented by the provided function.
func Filter[T any](arr []T, test func(T) bool) []T {
	var result []T
	for _, v := range arr {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

// Find returns the first element in the array that satisfies the provided testing function. If no values satisfy the testing function, the function returns the zero value for type T.
func Find[T any](arr []T, test func(T) bool) T {
	for _, v := range arr {
		if test(v) {
			return v
		}
	}
	var zero T
	return zero
}

// FindIndex returns the index of the first element in the array that satisfies the provided testing function. Otherwise, it returns -1, indicating that no element passed the test.
func FindIndex[T any](arr []T, test func(T) bool) int {
	for i, v := range arr {
		if test(v) {
			return i
		}
	}
	return -1
}

// Join creates and returns a new string by concatenating all of the elements in an array (or an array-like object), separated by commas or a specified separator string.
func Join[T any](arr []T, separator string) string {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		strArr[i] = fmt.Sprint(v)
	}
	return strings.Join(strArr, separator)
}

// FindLast returns the last element in the array that satisfies the provided testing function. If no values satisfy the testing function, the function returns the zero value for type T.
func FindLast[T any](arr []T, test func(T) bool) T {
	for i := len(arr) - 1; i >= 0; i-- {
		if test(arr[i]) {
			return arr[i]
		}
	}
	var zero T
	return zero
}

// FindLastIndex returns the index of the last element in the array that satisfies the provided testing function. Otherwise, it returns -1, indicating that no element passed the test.
func FindLastIndex[T any](arr []T, test func(T) bool) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if test(arr[i]) {
			return i
		}
	}
	return -1
}

// Flat returns a new array with all sub-array elements concatenated into it recursively up to the specified depth.
func Flat[T any](arr [][]T) []T {
	var result []T
	for _, v := range arr {
		result = append(result, v...)
	}
	return result
}

// FlatMap first maps each element using a mapping function, then flattens the result into a new array.
func FlatMap[T any, U any](arr []T, mapFunc func(T) []U) []U {
	var result []U
	for _, v := range arr {
		result = append(result, mapFunc(v)...)
	}
	return result
}

// ForEach executes a provided function once for each array element.
func ForEach[T any](arr []T, callback func(T)) {
	for _, v := range arr {
		callback(v)
	}
}

// Includes determines whether an array includes a certain value among its entries, returning true or false as appropriate.
func Includes[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// IndexOf returns the first index at which a given element can be found in the array, or -1 if it is not present.
func IndexOf[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

// Keys returns an array of a given object's own enumerable property names, iterated in the same order that a normal loop would.
func Keys[T any](arr []T) []int {
	keys := make([]int, len(arr))
	for i := range arr {
		keys[i] = i
	}
	return keys
}

// LastIndexOf returns the last index at which a given element can be found in the array, or -1 if it is not present.
func LastIndexOf[T comparable](arr []T, value T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == value {
			return i
		}
	}
	return -1
}

// Map creates a new array populated with the results of calling a provided function on every element in the calling array.
func Map[T any, U any](arr []T, mapFunc func(T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = mapFunc(v)
	}
	return result
}

// Pop removes the last element from an array and returns that element. This method changes the length of the array.
func Pop[T any](arr *[]T) T {
	if len(*arr) == 0 {
		var zero T
		return zero
	}
	lastElement := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return lastElement
}

// Push adds one or more elements to the end of an array and returns the new length of the array.
func Push[T any](arr *[]T, elements ...T) int {
	*arr = append(*arr, elements...)
	return len(*arr)
}

// Reduce executes a reducer function (that you provide) on each element of the array, resulting in a single output value.
func Reduce[T any, U any](arr []T, reducer func(U, T) U, initialValue U) U {
	result := initialValue
	for _, v := range arr {
		result = reducer(result, v)
	}
	return result
}

// ReduceRight applies a function against an accumulator and each value of the array (from right-to-left) to reduce it to a single value.
func ReduceRight[T any, U any](arr []T, reducer func(U, T) U, initialValue U) U {
	result := initialValue
	for i := len(arr) - 1; i >= 0; i-- {
		result = reducer(result, arr[i])
	}
	return result
}

// Reverse reverses an array in place.
func Reverse[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// Shift removes the first element from an array and returns that removed element. This method changes the length of the array.
func Shift[T any](arr *[]T) T {
	if len(*arr) == 0 {
		var zero T
		return zero
	}
	firstElement := (*arr)[0]
	*arr = (*arr)[1:]
	return firstElement
}

// Slice returns a shallow copy of a portion of an array into a new array object selected from start to end (end not included) where start and end represent the index of items in that array.
func Slice[T any](arr []T, start, end int) []T {
	if start < 0 {
		start = len(arr) + start
	}
	if end < 0 {
		end = len(arr) + end
	}
	if start < 0 {
		start = 0
	}
	if end > len(arr) {
		end = len(arr)
	}
	if start > end {
		return []T{}
	}
	return arr[start:end]
}

// Some tests whether at least one element in the array passes the test implemented by the provided function.
func Some[T any](arr []T, test func(T) bool) bool {
	for _, v := range arr {
		if test(v) {
			return true
		}
	}
	return false
}

// Sort sorts the elements of an array in place and returns the array. The default sort order is ascending.
func Sort[T any](arr []T, less func(T, T) bool) []T {
	sort.Slice(arr, func(i, j int) bool {
		return less(arr[i], arr[j])
	})
	return arr
}

// Splice changes the contents of an array by removing or replacing existing elements and/or adding new elements in place.
func Splice[T any](arr *[]T, start, deleteCount int, items ...T) []T {
	if start < 0 {
		start = len(*arr) + start
	}
	if start < 0 {
		start = 0
	} else if start > len(*arr) {
		start = len(*arr)
	}

	if deleteCount < 0 {
		deleteCount = 0
	} else if deleteCount > len(*arr)-start {
		deleteCount = len(*arr) - start
	}

	removed := append([]T(nil), (*arr)[start:start+deleteCount]...)
	*arr = append((*arr)[:start], append(items, (*arr)[start+deleteCount:]...)...)

	return removed
}

// ToReversed returns a new array with the elements in reversed order.
func ToReversed[T any](arr []T) []T {
	result := make([]T, len(arr))
	copy(result, arr)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// ToSorted returns a new array with the elements sorted in ascending order.
func ToSorted[T any](arr []T, less func(T, T) bool) []T {
	result := make([]T, len(arr))
	copy(result, arr)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

// ToSpliced returns a new array with the specified elements spliced.
func ToSpliced[T any](arr []T, start, deleteCount int, items ...T) []T {
	if start < 0 {
		start = len(arr) + start
	}
	if start < 0 {
		start = 0
	} else if start > len(arr) {
		start = len(arr)
	}

	if deleteCount < 0 {
		deleteCount = 0
	} else if deleteCount > len(arr)-start {
		deleteCount = len(arr) - start
	}

	result := append([]T(nil), arr[:start]...)
	result = append(result, items...)
	result = append(result, arr[start+deleteCount:]...)
	return result
}

// ToString returns a string representing the elements of the array.
func ToString[T any](arr []T) string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, fmt.Sprint(v))
	}
	return strings.Join(strArr, ",")
}

// Unshift adds one or more elements to the beginning of an array and returns the new length of the array.
func Unshift[T any](arr *[]T, elements ...T) int {
	*arr = append(elements, *arr...)
	return len(*arr)
}

// Values returns a new array iterator object that contains the values for each index in the array.
func Values[T any](arr []T) []T {
	return arr
}

// With returns a new array with the specified value inserted at the specified index.
func With[T any](arr []T, index int, value T) []T {
	result := append([]T(nil), arr...)
	if index >= 0 && index < len(arr) {
		result[index] = value
	}
	return result
}

type KeyValuePair[T any] struct {
	Key   int
	Value T
}

// Entries returns a slice of KeyValuePair containing the key/value pairs for each index in the array.
func Entries[T any](arr []T) []KeyValuePair[T] {
	entries := make([]KeyValuePair[T], len(arr))
	for i, v := range arr {
		entries[i] = KeyValuePair[T]{Key: i, Value: v}
	}
	return entries
}
