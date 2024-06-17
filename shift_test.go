package funky

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 1
	if got := Shift(&arr); got != want {
		t.Errorf("Shift() = %v, want %v", got, want)
	}
	wantArr := []int{2, 3, 4, 5}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Shift() arr = %v, want %v", arr, wantArr)
	}
}
