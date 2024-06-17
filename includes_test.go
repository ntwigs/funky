package funky

import "testing"

func TestIncludes(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if got := Includes(arr, 3); !got {
		t.Errorf("Includes() = %v, want %v", got, true)
	}
	if got := Includes(arr, 6); got {
		t.Errorf("Includes() = %v, want %v", got, false)
	}
}
