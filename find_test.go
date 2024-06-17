package funky

import "testing"

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := Find(arr, func(n int) bool { return n > 3 }); got != 4 {
		t.Errorf("Find() = %v, want %v", got, 4)
	}
}
