package funky

import "testing"

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 15
	if got := Reduce(arr, func(acc, n int) int { return acc + n }, 0); got != want {
		t.Errorf("Reduce() = %v, want %v", got, want)
	}
}
