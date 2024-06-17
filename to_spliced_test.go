package funky

import (
	"reflect"
	"testing"
)

func TestToSpliced(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := []int{1, 8, 9, 4, 5}
	if got := ToSpliced(arr, 1, 2, 8, 9); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSpliced() = %v, want %v", got, want)
	}
}
