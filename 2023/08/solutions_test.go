package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

//go:embed input_b_test.txt
var testBInput string

func TestPart1(t *testing.T) {
	correct := "6"
	ans := part1(testInput)
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart2(t *testing.T) {
	correct := "6"
	ans := part2(testBInput)
	if ans != correct {
		t.Errorf("part2() got %s; want %s", ans, correct)
	}
}
