package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

//go:embed input_test_b.txt
var testInputB string

func TestPart1(t *testing.T) {
	correct := "161"
	ans := part1(testInput)
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart2(t *testing.T) {
	correct := "48"
	ans := part2(testInputB)
	if ans != correct {
		t.Errorf("part2() got %s; want %s", ans, correct)
	}
}
