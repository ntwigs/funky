package funky

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	obj := map[string]int{"a": 1, "b": 2}
	want := []string{"a", "b"}

	result := Keys(obj)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Keys() = %v, want %v", result, want)
	}
}
