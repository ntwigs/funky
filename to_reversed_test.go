package funky

import (
	"reflect"
	"testing"
)

func TestToReversed(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{5, 4, 3, 2, 1}
	if got := ToReversed(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("ToReversed() = %v, want %v", got, want)
	}
}
