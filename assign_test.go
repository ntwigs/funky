package funky

import (
	"reflect"
	"testing"
)

func TestAssign(t *testing.T) {
	target := map[string]int{"a": 1}
	source1 := map[string]int{"b": 2}
	source2 := map[string]int{"c": 3}
	want := map[string]int{"a": 1, "b": 2, "c": 3}

	result := Assign(target, source1, source2)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Assign() = %v, want %v", result, want)
	}
}
