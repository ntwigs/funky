package funky

import "testing"

func TestReduceRight(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 15
	if got := ReduceRight(arr, func(acc, n int) int { return acc + n }, 0); got != want {
		t.Errorf("ReduceRight() = %v, want %v", got, want)
	}
}
