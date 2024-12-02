package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

func TestPart1(t *testing.T) {
	correct := "[ANSWER]"
	ans := part1(testInput)
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart2(t *testing.T) {
	correct := "[ANSWER]"
	ans := part2(testInput)
	if ans != correct {
		t.Errorf("part2() got %s; want %s", ans, correct)
	}
}

func TestIncrementPassword(t *testing.T) {
	correct := "xy"
	ans := incrementPassword("xx")
	if ans != correct {
		t.Errorf("incrementPassword() got %s; want %s", ans, correct)
	}
}

func TestIncrementPasswordWrapping(t *testing.T) {
	correct := "za"
	ans := incrementPassword("yz")
	if ans != correct {
		t.Errorf("incrementPassword() got %s; want %s", ans, correct)
	}
}

func TestIncrementPasswordWrappingLong(t *testing.T) {
	correct := "shfaoishaaa"
	ans := incrementPassword("shfaoisgzzz")
	if ans != correct {
		t.Errorf("incrementPassword() got %s; want %s", ans, correct)
	}
}
