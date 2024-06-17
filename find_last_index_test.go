package funky

import "testing"

func TestFindLastIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := FindLastIndex(arr, func(n int) bool { return n < 4 }); got != 2 {
		t.Errorf("FindLastIndex() = %v, want %v", got, 2)
	}
}
