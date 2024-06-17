package funky

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{2, 4}
	if got := Filter(arr, func(n int) bool { return n%2 == 0 }); !reflect.DeepEqual(got, want) {
		t.Errorf("Filter() = %v, want %v", got, want)
	}
}
