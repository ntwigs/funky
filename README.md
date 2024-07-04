<p align="center">
  <img alt='funky-sloth' src='https://github.com/ntwigs/funky/assets/14088342/1710632c-da0e-4575-8668-e374c617e3ce' width='250'/>
  <h1 align="center">Funky</h1>
  <p align="center">🔥 Utility functions for Go 🔥</p>
  <p align="center">
    <a href="https://pkg.go.dev/github.com/ntwigs/funky?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/ntwigs/funky/actions"><img src="https://github.com/ntwigs/funky/workflows/tests/badge.svg" alt="Build Status"></a>
    <a href="https://goreportcard.com/report/ntwigs/funky"><img src="https://goreportcard.com/badge/ntwigs/funky" alt="Go ReportCard"></a>
  </p>
</div>

---

A delightful little collection of Go utility functions. I think they're neat. 🥃

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

### With
Returns a new array with the specified value inserted at the specified index.

```go
arr := []int{1, 2, 3, 4, 5}
fmt.Println(funky.With(arr, 2, 10)) // Output: [1 2 10 4 5]
```

### Assign
Merges multiple source maps into a target map. Properties from the sources will overwrite those in the target map if they have the same key.

```go
obj1 := map[string]int{"a": 1, "b": 2}
obj2 := map[string]int{"b": 3, "c": 4}
fmt.Println(funky.Assign(obj1, obj2)) // Output: map[a:1 b:3 c:4]
```

### Has
Checks if a map contains a specific key. Returns true if the key exists, false otherwise.

```go
obj := map[string]int{"a": 1, "b": 2}
fmt.Println(funky.Has(obj, "a")) // Output: true
fmt.Println(funky.Has(obj, "dragon")) // Output: false
```

### Keys
Returns a slice containing all the keys in a map.

```go
obj := map[string]int{"a": 1, "b": 2}
fmt.Println(funky.Keys(obj)) // Output: [a b]
```

### Values
Returns a slice containing all the values in a map.

```go
obj := map[string]int{"a": 1, "b": 2}
fmt.Println(funky.Values(obj)) // Output: [1 2]
```

## Running Tests
Run the tests using the make command:

```sh
make test
```

---

Created with an unhealthy amount of hatred towards for-loops, mutations, and switch statements.
