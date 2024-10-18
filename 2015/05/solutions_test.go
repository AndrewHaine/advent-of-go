package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

func TestPart1(t *testing.T) {
	correct := "2"
	ans := part1(testInput)
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

//go:embed input_b_test.txt
var testInputB string

func TestPart2(t *testing.T) {
	correct := "2"
	ans := part2(testInputB)
	if ans != correct {
		t.Errorf("part2() got %s; want %s", ans, correct)
	}
}

