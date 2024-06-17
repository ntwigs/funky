package funky

import "testing"

func TestToString(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := "1,2,3,4,5"
	if got := ToString(arr); got != want {
		t.Errorf("ToString() = %v, want %v", got, want)
	}
}
