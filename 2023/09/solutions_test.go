package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

func TestPart1(t *testing.T) {
	correct := "114"
	ans := part1(testInput)
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

// func TestPart2(t *testing.T) {
// 	correct := "[ANSWER]"
// 	ans := part2(testInput)
// 	if ans != correct {
// 		t.Errorf("part2() got %s; want %s", ans, correct)
// 	}
// }
