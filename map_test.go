package funky

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	want := []int{2, 4, 6}
	if got := Map(arr, func(n int) int { return n * 2 }); !reflect.DeepEqual(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}
