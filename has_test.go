package funky

import "testing"

func TestHas(t *testing.T) {
	obj := map[string]int{"a": 1}
	key := "a"

	if !Has(obj, key) {
		t.Errorf("HasOwn() = false, want true")
	}
}
