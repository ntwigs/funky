package funky

import (
	"slices"
	"testing"
)

func TestValues(t *testing.T) {
	obj := map[string]int{"a": 1, "b": 2}
	want := []int{1, 2}

	result := Values(obj)

	for _, value := range want {
		hasValue := slices.Contains(result, value)

		if !hasValue {
			t.Errorf("Values() = %v, want %v", result, want)
		}
	}
}
