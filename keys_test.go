package funky

import (
	"slices"
	"testing"
)

func TestKeys(t *testing.T) {
	obj := map[string]int{"a": 1, "b": 2}
	want := []string{"a", "b"}

	result := Keys(obj)

	for _, key := range want {
		hasKey := slices.Contains(result, key)

		if !hasKey {
			t.Errorf("Keys() = %v, want %v", result, want)
		}
	}
}
