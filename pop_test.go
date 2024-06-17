package funky

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 5
	if got := Pop(&arr); got != want {
		t.Errorf("Pop() = %v, want %v", got, want)
	}
	wantArr := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Pop() arr = %v, want %v", arr, wantArr)
	}
}
