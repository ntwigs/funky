package funky

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{5, 3, 1, 4, 2}
	want := []int{1, 2, 3, 4, 5}
	if got := Sort(arr, func(a, b int) bool { return a < b }); !reflect.DeepEqual(got, want) {
		t.Errorf("Sort() = %v, want %v", got, want)
	}
}
