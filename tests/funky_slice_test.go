package tests

import (
	"reflect"
	"testing"

	"github.com/ntwigs/funky"
)

func TestAt(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	if got := funky.At(arr, 2); got != 3 {
		t.Errorf("At() = %v, want %v", got, 3)
	}
	if got := funky.At(arr, -1); got != 5 {
		t.Errorf("At() = %v, want %v", got, 5)
	}
	if got := funky.At(arr, 5); got != 0 {
		t.Errorf("At() = %v, want %v", got, 0)
	}
}

func TestConcat(t *testing.T) {
	arr := []int{1, 2, 3}
	arr2 := []int{4, 5}
	want := []int{1, 2, 3, 4, 5}
	if got := funky.Concat(arr, arr2); !reflect.DeepEqual(got, want) {
		t.Errorf("Concat() = %v, want %v", got, want)
	}
}

func TestCopyWithin(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{4, 5, 3, 4, 5}
	if got := funky.CopyWithin(arr, 0, 3, 5); !reflect.DeepEqual(got, want) {
		t.Errorf("CopyWithin() = %v, want %v", got, want)
	}
}

func TestEvery(t *testing.T) {
	arr := []int{2, 4, 6}
	if got := funky.Every(arr, func(n int) bool { return n%2 == 0 }); !got {
		t.Errorf("Every() = %v, want %v", got, true)
	}
}

func TestFill(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{1, 0, 0, 0, 5}
	if got := funky.Fill(arr, 0, 1, 4); !reflect.DeepEqual(got, want) {
		t.Errorf("Fill() = %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{2, 4}
	if got := funky.Filter(arr, func(n int) bool { return n%2 == 0 }); !reflect.DeepEqual(got, want) {
		t.Errorf("Filter() = %v, want %v", got, want)
	}
}

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := funky.Find(arr, func(n int) bool { return n > 3 }); got != 4 {
		t.Errorf("Find() = %v, want %v", got, 4)
	}
}

func TestFindIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := funky.FindIndex(arr, func(n int) bool { return n > 3 }); got != 3 {
		t.Errorf("FindIndex() = %v, want %v", got, 3)
	}
}

func TestFindLast(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := funky.FindLast(arr, func(n int) bool { return n < 4 }); got != 3 {
		t.Errorf("FindLast() = %v, want %v", got, 3)
	}
}

func TestFindLastIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := funky.FindLastIndex(arr, func(n int) bool { return n < 4 }); got != 2 {
		t.Errorf("FindLastIndex() = %v, want %v", got, 2)
	}
}

func TestFlat(t *testing.T) {
	arr := [][]int{{1, 2}, {3, 4}, {5}}
	want := []int{1, 2, 3, 4, 5}
	if got := funky.Flat(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("Flat() = %v, want %v", got, want)
	}
}

func TestFlatMap(t *testing.T) {
	arr := []int{1, 2, 3}
	want := []int{1, -1, 2, -2, 3, -3}
	if got := funky.FlatMap(arr, func(n int) []int { return []int{n, -n} }); !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}

func TestForEach(t *testing.T) {
	arr := []int{1, 2, 3}
	sum := 0
	funky.ForEach(arr, func(n int) { sum += n })
	if sum != 6 {
		t.Errorf("ForEach() sum = %v, want %v", sum, 6)
	}
}

func TestIncludes(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := funky.Includes(arr, 3); !got {
		t.Errorf("Includes() = %v, want %v", got, true)
	}
	if got := funky.Includes(arr, 6); got {
		t.Errorf("Includes() = %v, want %v", got, false)
	}
}

func TestIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := funky.IndexOf(arr, 3); got != 2 {
		t.Errorf("IndexOf() = %v, want %v", got, 2)
	}
	if got := funky.IndexOf(arr, 6); got != -1 {
		t.Errorf("IndexOf() = %v, want %v", got, -1)
	}
}

func TestJoin(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := "1-2-3-4-5"
	if got := funky.Join(arr, "-"); got != want {
		t.Errorf("Join() = %v, want %v", got, want)
	}
}

func TestLastIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 2, 5}
	if got := funky.LastIndexOf(arr, 2); got != 3 {
		t.Errorf("LastIndexOf() = %v, want %v", got, 3)
	}
	if got := funky.LastIndexOf(arr, 6); got != -1 {
		t.Errorf("LastIndexOf() = %v, want %v", got, -1)
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	want := []int{2, 4, 6}
	if got := funky.Map(arr, func(n int) int { return n * 2 }); !reflect.DeepEqual(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestPop(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 5
	if got := funky.Pop(&arr); got != want {
		t.Errorf("Pop() = %v, want %v", got, want)
	}
	wantArr := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Pop() arr = %v, want %v", arr, wantArr)
	}
}

func TestPush(t *testing.T) {
	arr := []int{1, 2, 3}
	wantLength := 5
	if got := funky.Push(&arr, 4, 5); got != wantLength {
		t.Errorf("Push() = %v, want %v", got, wantLength)
	}
	wantArr := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Push() arr = %v, want %v", arr, wantArr)
	}
}

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 15
	if got := funky.Reduce(arr, func(acc, n int) int { return acc + n }, 0); got != want {
		t.Errorf("Reduce() = %v, want %v", got, want)
	}
}

func TestReduceRight(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 15
	if got := funky.ReduceRight(arr, func(acc, n int) int { return acc + n }, 0); got != want {
		t.Errorf("ReduceRight() = %v, want %v", got, want)
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{5, 4, 3, 2, 1}
	if got := funky.Reverse(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("Reverse() = %v, want %v", got, want)
	}
}

func TestShift(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 1
	if got := funky.Shift(&arr); got != want {
		t.Errorf("Shift() = %v, want %v", got, want)
	}
	wantArr := []int{2, 3, 4, 5}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Shift() arr = %v, want %v", arr, wantArr)
	}
}

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{2, 3, 4}
	if got := funky.Slice(arr, 1, 4); !reflect.DeepEqual(got, want) {
		t.Errorf("Slice() = %v, want %v", got, want)
	}
}

func TestSome(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := funky.Some(arr, func(n int) bool { return n > 3 }); !got {
		t.Errorf("Some() = %v, want %v", got, true)
	}
	if got := funky.Some(arr, func(n int) bool { return n > 5 }); got {
		t.Errorf("Some() = %v, want %v", got, false)
	}
}

func TestSort(t *testing.T) {
	arr := []int{5, 3, 1, 4, 2}
	want := []int{1, 2, 3, 4, 5}
	if got := funky.Sort(arr, func(a, b int) bool { return a < b }); !reflect.DeepEqual(got, want) {
		t.Errorf("Sort() = %v, want %v", got, want)
	}
}

func TestSplice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{3}
	if got := funky.Splice(&arr, 2, 1, 6, 7); !reflect.DeepEqual(got, want) {
		t.Errorf("Splice() = %v, want %v", got, want)
	}
	wantArr := []int{1, 2, 6, 7, 4, 5}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Splice() arr = %v, want %v", arr, wantArr)
	}
}

func TestToReversed(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{5, 4, 3, 2, 1}
	if got := funky.ToReversed(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("ToReversed() = %v, want %v", got, want)
	}
}

func TestToSorted(t *testing.T) {
	arr := []int{5, 3, 1, 4, 2}
	want := []int{1, 2, 3, 4, 5}
	if got := funky.ToSorted(arr, func(a, b int) bool { return a < b }); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSorted() = %v, want %v", got, want)
	}
}

func TestToSpliced(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{1, 8, 9, 4, 5}
	if got := funky.ToSpliced(arr, 1, 2, 8, 9); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSpliced() = %v, want %v", got, want)
	}
}

func TestToString(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := "1,2,3,4,5"
	if got := funky.ToString(arr); got != want {
		t.Errorf("ToString() = %v, want %v", got, want)
	}
}

func TestUnshift(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	wantLength := 6
	if got := funky.Unshift(&arr, 0); got != wantLength {
		t.Errorf("Unshift() = %v, want %v", got, wantLength)
	}
	wantArr := []int{0, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Unshift() arr = %v, want %v", arr, wantArr)
	}
}

func TestWith(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 10, 4, 5}
	if got := funky.With(arr, 2, 10); !reflect.DeepEqual(got, want) {
		t.Errorf("With() = %v, want %v", got, want)
	}
}
