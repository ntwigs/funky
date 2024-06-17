package funky

import "testing"

func TestForEach(t *testing.T) {
	arr := []int{1, 2, 3}
	sum := 0
	ForEach(arr, func(n int) { sum += n })
	if sum != 6 {
		t.Errorf("ForEach() sum = %v, want %v", sum, 6)
	}
}
