package funky

import "testing"

func TestSome(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := Some(arr, func(n int) bool { return n > 3 }); !got {
		t.Errorf("Some() = %v, want %v", got, true)
	}
	if got := Some(arr, func(n int) bool { return n > 5 }); got {
		t.Errorf("Some() = %v, want %v", got, false)
	}
}
