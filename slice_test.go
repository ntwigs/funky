package funky

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{2, 3, 4}
	if got := Slice(arr, 1, 4); !reflect.DeepEqual(got, want) {
		t.Errorf("Slice() = %v, want %v", got, want)
	}
}
