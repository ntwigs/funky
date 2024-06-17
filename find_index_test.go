package funky

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := FindIndex(arr, func(n int) bool { return n > 3 }); got != 3 {
		t.Errorf("FindIndex() = %v, want %v", got, 3)
	}
}
