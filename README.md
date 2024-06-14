# funky

The funky package provides a collection of array utility functions in Go, mimicking JavaScript/TypeScript Array methods.

I've created this, what some (most) people might call, **monstrosity**, since I am seriously missing all the array methods from JavaScript/TypeScript. Most of all, I missed my *map*, *filter* and *reduce*. Love those badbois. But I thought, might as well implement most array methods while I'm at it.

⚠️ Not battle-tested whatsoever - beware - danger ahead ⚠️

---

## Installation

```sh
go get github.com/ntwigs/funky
```

## Importing

```go
import "github.com/ntwigs/funky"
```

## Functions
### At
Returns the element at the specified index.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.At(arr, 2))  // Output: 3
fmt.Println(funky.At(arr, -1)) // Output: 5
```

### Concat
Returns a new array that is the result of concatenating the provided arrays.

```go
arr1 := []int{1, 2, 3}
arr2 := []int{4, 5}
fmt.Println(funky.Concat(arr1, arr2)) // Output: [1 2 3 4 5]
```

### CopyWithin
Shallow copies part of an array to another location in the same array and returns it without modifying its length.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.CopyWithin(arr, 0, 3, 5)) // Output: [4 5 3 4 5]
```

### Every
Tests whether all elements in the array pass the test implemented by the provided function.

```go
arr := []int{2, 4, 6}
fmt.Println(funky.Every(arr, func(n int) bool { return n%2 == 0 })) // Output: true
```

### Fill
Changes all elements in an array to a static value, from a start index to an end index.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Fill(arr, 0, 1, 4)) // Output: [1 0 0 0 5]
```

### Filter
Creates a new array with all elements that pass the test implemented by the provided function.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Filter(arr, func(n int) bool { return n%2 == 0 })) // Output: [2 4]
```

### Find
Returns the first element in the array that satisfies the provided testing function.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Find(arr, func(n int) bool { return n > 3 })) // Output: 4
```

### FindIndex
Returns the index of the first element in the array that satisfies the provided testing function.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.FindIndex(arr, func(n int) bool { return n > 3 })) // Output: 3
```

### FindLast
Returns the last element in the array that satisfies the provided testing function.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.FindLast(arr, func(n int) bool { return n < 4 })) // Output: 3
```

### FindLastIndex
Returns the index of the last element in the array that satisfies the provided testing function.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.FindLastIndex(arr, func(n int) bool { return n < 4 })) // Output: 2
```

### Flat
Returns a new array with all sub-array elements concatenated into it.

```go
arr := [][]int{{1, 2}, {3, 4}, {5}}
fmt.Println(funky.Flat(arr)) // Output: [1 2 3 4 5]
```

### FlatMap
First maps each element using a mapping function, then flattens the result into a new array.

```go
arr := []int{1, 2, 3}
fmt.Println(funky.FlatMap(arr, func(n int) []int { return []int{n, -n} })) // Output: [1 -1 2 -2 3 -3]
```

### ForEach
Executes a provided function once for each array element.

```go
arr := []int{1, 2, 3}
sum := 0
funky.ForEach(arr, func(n int) { sum += n })
fmt.Println(sum) // Output: 6
```

### Includes
Determines whether an array includes a certain value among its entries.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Includes(arr, 3)) // Output: true
fmt.Println(funky.Includes(arr, 6)) // Output: false
```

### IndexOf
Returns the first index at which a given element can be found in the array.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.IndexOf(arr, 3)) // Output: 2
fmt.Println(funky.IndexOf(arr, 6)) // Output: -1
```

### Join
Creates and returns a new string by concatenating all of the elements in an array.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Join(arr, "-")) // Output: "1-2-3-4-5"
```

### Keys
Returns an array of a given object's own enumerable property names.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Keys(arr)) // Output: [0 1 2 3 4]
```

### LastIndexOf
Returns the last index at which a given element can be found in the array.

```go
arr := []int{1, 2, 3, 2, 5}
fmt.Println(funky.LastIndexOf(arr, 2)) // Output: 3
```

### Map
Creates a new array populated with the results of calling a provided function on every element in the calling array.

```go
arr := []int{1, 2, 3}
fmt.Println(funky.Map(arr, func(n int) int { return n * 2 })) // Output: [2 4 6]
```
 
### Pop
Removes the last element from an array and returns that element.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Pop(&arr)) // Output: 5
fmt.Println(arr)             // Output: [1 2 3 4]
```

### Push
Adds one or more elements to the end of an array and returns the new length of the array.

```go
arr := []int{1, 2, 3}
fmt.Println(funky.Push(&arr, 4, 5)) // Output: 5
fmt.Println(arr)                    // Output: [1 2 3 4 5]
```

### Reduce
Executes a reducer function on each element of the array, resulting in a single output value.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Reduce(arr, func(acc, n int) int { return acc + n }, 0)) // Output: 15
```

### ReduceRight
Applies a function against an accumulator and each value of the array (from right-to-left) to reduce it to a single value.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.ReduceRight(arr, func(acc, n int) int { return acc + n }, 0)) // Output: 15
```

### Reverse
Reverses an array in place.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Reverse(arr)) // Output: [5 4 3 2 1]
```

### Shift
Removes the first element from an array and returns that removed element.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Shift(&arr)) // Output: 1
fmt.Println(arr)               // Output: [2 3 4 5]
```

### Slice
Returns a shallow copy of a portion of an array into a new array object selected from start to end (end not included).

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Slice(arr, 1, 4)) // Output: [2 3 4]
```

### Some
Tests whether at least one element in the array passes the test implemented by the provided function.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Some(arr, func(n int) bool { return n > 3 })) // Output: true
```

### Sort
Sorts the elements of an array in place and returns the array.

```go
arr := []int{5, 3, 1, 4, 2}
fmt.Println(funky.Sort(arr, func(a, b int) bool { return a < b })) // Output: [1 2 3 4 5]
```

### Splice
Changes the contents of an array by removing or replacing existing elements and/or adding new elements in place.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Splice(&arr, 2, 1, 6, 7)) // Output: [3]
fmt.Println(arr)                           // Output: [1 2 6 7 4 5]
```

### ToReversed
Returns a new array with the elements in reversed order.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.ToReversed(arr)) // Output: [5 4 3 2 1]
```

### ToSorted
Returns a new array with the elements sorted in ascending order.

```go
arr := []int{5, 3, 1, 4, 2}
fmt.Println(funky.ToSorted(arr, func(a, b int) bool { return a < b })) // Output: [1 2 3 4 5]
```

### ToSpliced
Returns a new array with the specified elements spliced.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.ToSpliced(arr, 1, 2, 8, 9)) // Output: [1 8 9 4 5]
```

### ToString
Returns a string representing the elements of the array.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.ToString(arr)) // Output: "1,2,3,4,5"
```

### Unshift
Adds one or more elements to the beginning of an array and returns the new length of the array.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Unshift(&arr, 0)) // Output: 6
fmt.Println(arr)                    // Output: [0 1 2 3 4 5]
```

### Values
Returns a new array iterator object that contains the values for each index in the array.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.Values(arr)) // Output: [1 2 3 4 5]
```

### With
Returns a new array with the specified value inserted at the specified index.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.With(arr, 2, 10)) // Output: [1 2 10 4 5]
```

### Entries
Returns a slice of KeyValuePair containing the key/value pairs for each index in the array.

```go
arr := []int{1, 2, 3}
entries := funky.Entries(arr)
for _, entry := range entries {
	fmt.Printf("Key: %d, Value: %d\n", entry.Key, entry.Value)
}
// Output:
// Key: 0, Value: 1
// Key: 1, Value: 2
// Key: 2, Value: 3
```

## Running Tests
Run the tests using the make command:

```sh
make test
```

---

License
This project is licensed under the MIT License.

