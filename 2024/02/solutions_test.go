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

func TestPart2(t *testing.T) {
	correct := "4"
	ans := part2(testInput)
	if ans != correct {
		t.Errorf("part2() got %s; want %s", ans, correct)
	}
}

func TestIsReportSafeSafe(t *testing.T) {
	correct := true
	ans := isReportSafe("7 6 4 2 1")
	if ans != correct {
		t.Errorf("isReportSafe() got %t; want %t", ans, correct)
	}
}

func TestIsReportSafeUnsafeDirectionChange(t *testing.T) {
	correct := false
	ans := isReportSafe("1 3 2 4 5")
	if ans != correct {
		t.Errorf("isReportSafe() got %t; want %t", ans, correct)
	}
}

func TestIsReportSafeUnsafeNoChange(t *testing.T) {
	correct := false
	ans := isReportSafe("8 6 4 4 1")
	if ans != correct {
		t.Errorf("isReportSafe() got %t; want %t", ans, correct)
	}
}

func TestIsReportSafeUnsafeLargeDecrease(t *testing.T) {
	correct := false
	ans := isReportSafe("9 7 6 2 1")
	if ans != correct {
		t.Errorf("isReportSafe() got %t; want %t", ans, correct)
	}
}

func TestIsReportSafeUnsafeLargeIncrease(t *testing.T) {
	correct := false
	ans := isReportSafe("1 2 7 8 9")
	if ans != correct {
		t.Errorf("isReportSafe() got %t; want %t", ans, correct)
	}
}
