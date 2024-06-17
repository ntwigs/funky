package funky

import (
	"reflect"
	"testing"
)

func TestUnshift(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	wantLength := 6
	if got := Unshift(&arr, 0); got != wantLength {
		t.Errorf("Unshift() = %v, want %v", got, wantLength)
	}
	wantArr := []int{0, 1, 2, 3, 4, 5}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Unshift() arr = %v, want %v", arr, wantArr)
	}
}
