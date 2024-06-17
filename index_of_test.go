package funky

import "testing"

func TestIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := IndexOf(arr, 3); got != 2 {
		t.Errorf("IndexOf() = %v, want %v", got, 2)
	}
	if got := IndexOf(arr, 6); got != -1 {
		t.Errorf("IndexOf() = %v, want %v", got, -1)
	}
}
