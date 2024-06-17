package funky

import "testing"

func TestFindLast(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := FindLast(arr, func(n int) bool { return n < 4 }); got != 3 {
		t.Errorf("FindLast() = %v, want %v", got, 3)
	}
}
