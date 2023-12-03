package sliceutil

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	initial := []int{1, 2, 3, 4}
	correct := []int{1, 2, 3}

	_, newSlice := Pop(initial)

	if !reflect.DeepEqual(newSlice, correct) {
		t.Errorf("Pop got %v; want %v", newSlice, correct)
	}
}
