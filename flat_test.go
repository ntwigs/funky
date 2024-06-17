package funky

import (
	"reflect"
	"testing"
)

func TestFlat(t *testing.T) {
	arr := [][]int{{1, 2}, {3, 4}, {5}}
	want := []int{1, 2, 3, 4, 5}
	if got := Flat(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("Flat() = %v, want %v", got, want)
	}
}
