package funky

import (
	"reflect"
	"testing"
)

func TestToSorted(t *testing.T) {
	arr := []int{5, 3, 1, 4, 2}
	want := []int{1, 2, 3, 4, 5}
	if got := ToSorted(arr, func(a, b int) bool { return a < b }); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSorted() = %v, want %v", got, want)
	}
}
