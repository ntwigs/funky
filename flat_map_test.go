package funky

import (
	"reflect"
	"testing"
)

func TestFlatMap(t *testing.T) {
	arr := []int{1, 2, 3}
	want := []int{1, -1, 2, -2, 3, -3}
	if got := FlatMap(arr, func(n int) []int { return []int{n, -n} }); !reflect.DeepEqual(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}
