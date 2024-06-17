package funky

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	arr := []int{1, 2, 3}
	wantLength := 5
	if got := Push(&arr, 4, 5); got != wantLength {
		t.Errorf("Push() = %v, want %v", got, wantLength)
	}
	wantArr := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(arr, wantArr) {
		t.Errorf("Push() arr = %v, want %v", arr, wantArr)
	}
}
