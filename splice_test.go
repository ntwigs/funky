package funky

import (
	"reflect"
	"testing"
)

func TestSplice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{3}
	if got := Splice(&arr, 2, 1, 6, 7); !reflect.DeepEqual(got, want) {
		t.Errorf("Splice() = %v, want %v", got, want)
	}
	wantArr := []int{1, 2, 6, 7, 4, 5}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Splice() arr = %v, want %v", arr, wantArr)
	}
}
