package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

func TestPart1(t *testing.T) {
	correct := "12"
	ans := part1(testInput)
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestCodeCharLengthEmptyString(t *testing.T) {
	correct := 2
	ans := getCodeCharLength("\"\"")
	if ans != correct {
		t.Errorf("getCodeCharLength() for empty string got %d; want %d", ans, correct)
	}
}

func TestCodeCharLengthNormalString(t *testing.T) {
	correct := 5
	ans := getCodeCharLength("\"abc\"")
	if ans != correct {
		t.Errorf("getCodeCharLength() for \"abc\" got %d; want %d", ans, correct)
	}
}

func TestCodeCharLengthEscapedQuote(t *testing.T) {
	correct := 10
	ans := getCodeCharLength("\"aaa\\\"aaa\"")
	if ans != correct {
		t.Errorf("getCodeCharLength() for \"aaa\\\"aaa\" got %d; want %d", ans, correct)
	}
}

func TestCodeCharLengthEscapedHex(t *testing.T) {
	correct := 6
	ans := getCodeCharLength("\"\\x27\"")
	if ans != correct {
		t.Errorf("getCodeCharLength() for \"\\x27\" got %d; want %d", ans, correct)
	}
}

func TestMemoryCharLengthEmptyString(t *testing.T) {
	correct := 0
	ans := getMemoryCharLength("\"\"")
	if ans != correct {
		t.Errorf("getCodeCharLength() for empty string got %d; want %d", ans, correct)
	}
}

func TestMemoryCharLengthNormalString(t *testing.T) {
	correct := 3
	ans := getMemoryCharLength("\"abc\"")
	if ans != correct {
		t.Errorf("getCodeCharLength() for \"abc\" got %d; want %d", ans, correct)
	}
}

func TestMemoryCharLengthEscapedQuote(t *testing.T) {
	correct := 7
	ans := getMemoryCharLength("\"aaa\\\"aaa\"")
	if ans != correct {
		t.Errorf("getCodeCharLength() for \"aaa\\\"aaa\" got %d; want %d", ans, correct)
	}
}

func TestMemoryCharLengthEscapedHex(t *testing.T) {
	correct := 1
	ans := getMemoryCharLength("\"\\x27\"")
	if ans != correct {
		t.Errorf("getCodeCharLength() for \"\\x27\" got %d; want %d", ans, correct)
	}
}

func TestPart2(t *testing.T) {
	correct := "19"
	ans := part2(testInput)
	if ans != correct {
		t.Errorf("part2() got %s; want %s", ans, correct)
	}
}
