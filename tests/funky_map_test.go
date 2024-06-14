package tests

import (
	"reflect"
	"testing"

	"github.com/ntwigs/funky/funky"
)

func TestAssign(t *testing.T) {
	target := map[string]int{"a": 1}
	source1 := map[string]int{"b": 2}
	source2 := map[string]int{"c": 3}
	want := map[string]int{"a": 1, "b": 2, "c": 3}

	result := funky.Assign(target, source1, source2)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Assign() = %v, want %v", result, want)
	}
}

func TestHas(t *testing.T) {
	obj := map[string]int{"a": 1}
	key := "a"

	if !funky.Has(obj, key) {
		t.Errorf("HasOwn() = false, want true")
	}
}

func TestKeys(t *testing.T) {
	obj := map[string]int{"a": 1, "b": 2}
	want := []string{"a", "b"}

	result := funky.Keys(obj)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Keys() = %v, want %v", result, want)
	}
}

func TestValues(t *testing.T) {
	obj := map[string]int{"a": 1, "b": 2}
	want := []int{1, 2}

	result := funky.Values(obj)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Values() = %v, want %v", result, want)
	}
}
