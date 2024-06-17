package funky

import "testing"

func TestLastIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 2, 5}
	if got := LastIndexOf(arr, 2); got != 3 {
		t.Errorf("LastIndexOf() = %v, want %v", got, 3)
	}
	if got := LastIndexOf(arr, 6); got != -1 {
		t.Errorf("LastIndexOf() = %v, want %v", got, -1)
	}
}
