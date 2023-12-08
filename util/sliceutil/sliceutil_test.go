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

func TestEveryTrue(t *testing.T) {
	correct := true
	slice := []int{2, 4, 6, 8, 10}
	allEven := Every[int](slice, func(i int, item int) bool {
		return item%2 == 0
	})
	if !allEven {
		t.Errorf("Every got %v; want %v", allEven, correct)
	}
}

func TestEveryFalse(t *testing.T) {
	correct := false
	slice := []int{2, 4, 5, 8, 10}
	allEven := Every[int](slice, func(i int, item int) bool {
		return item%2 == 0
	})
	if allEven {
		t.Errorf("Every got %v; want %v", allEven, correct)
	}
}

func TestSomeTrue(t *testing.T) {
	correct := true
	slice := []int{2, 4, 5, 8, 10}
	someOdd := Some[int](slice, func(i int, item int) bool {
		return item%2 != 0
	})
	if !someOdd {
		t.Errorf("Some got %v; want %v", someOdd, correct)
	}
}

func TestSomeFalse(t *testing.T) {
	correct := false
	slice := []int{2, 4, 6, 8, 10}
	someOdd := Some[int](slice, func(i int, item int) bool {
		return item%2 != 0
	})
	if someOdd {
		t.Errorf("Some got %v; want %v", someOdd, correct)
	}
}
