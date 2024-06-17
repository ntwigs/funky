package funky

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	arr := []int{1, 2, 3}
	arr2 := []int{4, 5}
	want := []int{1, 2, 3, 4, 5}
	if got := Concat(arr, arr2); !reflect.DeepEqual(got, want) {
		t.Errorf("Concat() = %v, want %v", got, want)
	}
}
