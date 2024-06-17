package funky

import (
	"testing"
)

func TestJoin(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := "1-2-3-4-5"
	if got := Join(arr, "-"); got != want {
		t.Errorf("Join() = %v, want %v", got, want)
	}
}
