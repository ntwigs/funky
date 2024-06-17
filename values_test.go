package funky

import (
	"reflect"
	"testing"
)

func TestValues(t *testing.T) {
	obj := map[string]int{"a": 1, "b": 2}
	want := []int{1, 2}

	result := Values(obj)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Values() = %v, want %v", result, want)
	}
}
