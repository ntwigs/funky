package funky

import (
	"reflect"
	"testing"
)

func TestWith(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 10, 4, 5}
	if got := With(arr, 2, 10); !reflect.DeepEqual(got, want) {
		t.Errorf("With() = %v, want %v", got, want)
	}
}
