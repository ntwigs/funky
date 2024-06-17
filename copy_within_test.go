package funky

import (
	"reflect"
	"testing"
)

func TestCopyWithin(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{4, 5, 3, 4, 5}
	if got := CopyWithin(arr, 0, 3, 5); !reflect.DeepEqual(got, want) {
		t.Errorf("CopyWithin() = %v, want %v", got, want)
	}
}
