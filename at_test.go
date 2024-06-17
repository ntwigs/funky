package funky

import "testing"

func TestAt(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	if got := At(arr, 2); got != 3 {
		t.Errorf("At() = %v, want %v", got, 3)
	}
	if got := At(arr, -1); got != 5 {
		t.Errorf("At() = %v, want %v", got, 5)
	}
	if got := At(arr, 5); got != 0 {
		t.Errorf("At() = %v, want %v", got, 0)
	}
}
