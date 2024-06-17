package funky

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{5, 4, 3, 2, 1}
	if got := Reverse(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("Reverse() = %v, want %v", got, want)
	}
}
