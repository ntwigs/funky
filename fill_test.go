package funky

import (
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{1, 0, 0, 0, 5}
	if got := Fill(arr, 0, 1, 4); !reflect.DeepEqual(got, want) {
		t.Errorf("Fill() = %v, want %v", got, want)
	}
}
