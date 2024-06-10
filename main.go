package main

import (
	"fmt"

	"github.com/ntwigs/funky/funky"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}

	// Example usage of At
	fmt.Println(funky.At(arr, 2))  // Output: 3
	fmt.Println(funky.At(arr, -1)) // Output: 5

	// Example usage of Concat
	arr2 := []int{6, 7, 8}
	concatResult := funky.Concat(arr, arr2)
	fmt.Println(concatResult) // Output: [1 2 3 4 5 6 7 8]

	// Example usage of CopyWithin
	copyResult := funky.CopyWithin([]int{1, 2, 3, 4, 5}, 0, 3, 5)
	fmt.Println(copyResult) // Output: [4 5 3 4 5]

	// Example usage of Every
	isEveryPositive := funky.Every(arr, func(n int) bool { return n > 0 })
	fmt.Println(isEveryPositive) // Output: true

	// Example usage of Fill
	fillResult := funky.Fill([]int{1, 2, 3, 4, 5}, 0, 1, 4)
	fmt.Println(fillResult) // Output: [1 0 0 0 5]

	// Example usage of Filter
	filterResult := funky.Filter(arr, func(n int) bool { return n%2 == 0 })
	fmt.Println(filterResult) // Output: [2 4]

	// Example usage of Find
	findResult := funky.Find(arr, func(n int) bool { return n > 3 })
	fmt.Println(findResult) // Output: 4

	// Example usage of FindIndex
	findIndexResult := funky.FindIndex(arr, func(n int) bool { return n > 3 })
	fmt.Println(findIndexResult) // Output: 3

	// Example usage of Join
	joinResult := funky.Join(arr, "-")
	fmt.Println(joinResult) // Output: "1-2-3-4-5"

	// Example usage of LastIndexOf
	lastIndexResult := funky.LastIndexOf(arr, 3)
	fmt.Println(lastIndexResult) // Output: 2

	// Example usage of Map
	mapResult := funky.Map(arr, func(n int) int { return n * 2 })
	fmt.Println(mapResult) // Output: [2 4 6 8 10]

	// Example usage of Pop
	popResult := funky.Pop(&arr)
	fmt.Println(popResult) // Output: 5
	fmt.Println(arr)       // Output: [1 2 3 4]

	// Example usage of Push
	newLength := funky.Push(&arr, 5, 6)
	fmt.Println(newLength) // Output: 6
	fmt.Println(arr)       // Output: [1 2 3 4 5 6]

	// Example usage of Reduce
	reduceResult := funky.Reduce(arr, func(acc, n int) int { return acc + n }, 0)
	fmt.Println(reduceResult) // Output: 21

	// Example usage of ReduceRight
	reduceRightResult := funky.ReduceRight(arr, func(acc, n int) int { return acc + n }, 0)
	fmt.Println(reduceRightResult) // Output: 21

	// Example usage of Reverse
	reverseResult := funky.Reverse(arr)
	fmt.Println(reverseResult) // Output: [6 5 4 3 2 1]

	// Example usage of Shift
	shiftResult := funky.Shift(&arr)
	fmt.Println(shiftResult) // Output: 6
	fmt.Println(arr)         // Output: [5 4 3 2 1]

	// Example usage of Slice
	sliceResult := funky.Slice(arr, 1, 3)
	fmt.Println(sliceResult) // Output: [4 3]

	// Example usage of Some
	isSomeGreaterThanThree := funky.Some(arr, func(n int) bool { return n > 3 })
	fmt.Println(isSomeGreaterThanThree) // Output: true

	// Example usage of Sort
	sortResult := funky.Sort(arr, func(a, b int) bool { return a < b })
	fmt.Println(sortResult) // Output: [1 2 3 4 5]

	// Example usage of Splice
	arrToSplice := []int{1, 2, 3, 4, 5}
	spliceResult := funky.Splice(&arrToSplice, 2, 1, 6, 7)
	fmt.Println(spliceResult) // Output: [3]
	fmt.Println(arrToSplice)  // Output: [1 2 6 7 4 5]

	// Example usage of ToReversed
	toReversedResult := funky.ToReversed(arr)
	fmt.Println(toReversedResult) // Output: [5 4 3 2 1]

	// Example usage of ToSorted
	toSortedResult := funky.ToSorted(arr, func(a, b int) bool { return a > b })
	fmt.Println(toSortedResult) // Output: [5 4 3 2 1]

	// Example usage of ToSpliced
	toSplicedResult := funky.ToSpliced(arr, 1, 2, 8, 9)
	fmt.Println(toSplicedResult) // Output: [1 8 9 4 5]

	// Example usage of ToString
	toStringResult := funky.ToString(arr)
	fmt.Println(toStringResult) // Output: "1,2,3,4,5"

	// Example usage of Unshift
	unshiftResult := funky.Unshift(&arr, 0)
	fmt.Println(unshiftResult) // Output: 6
	fmt.Println(arr)           // Output: [0 1 2 3 4 5]

	// Example usage of Values
	valuesResult := funky.Values(arr)
	fmt.Println(valuesResult) // Output: [0 1 2 3 4 5]

	// Example usage of With
	withResult := funky.With(arr, 2, 10)
	fmt.Println(withResult) // Output: [0 1 10 3 4 5]

	// Example usage of Entries
	entriesResult := funky.Entries(arr)
	for _, entry := range entriesResult {
		fmt.Printf("Key: %d, Value: %d\n", entry.Key, entry.Value)
	}
}
