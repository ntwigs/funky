package funky

import "testing"

func TestEvery(t *testing.T) {
	arr := []int{2, 4, 6}
	if got := Every(arr, func(n int) bool { return n%2 == 0 }); !got {
		t.Errorf("Every() = %v, want %v", got, true)
	}
}
