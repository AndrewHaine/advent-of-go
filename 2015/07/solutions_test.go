package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

func TestPart1D(t *testing.T) {
	correct := "72"
	ans := part1(testInput, "d")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart1E(t *testing.T) {
	correct := "507"
	ans := part1(testInput, "e")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart1F(t *testing.T) {
	correct := "492"
	ans := part1(testInput, "f")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart1G(t *testing.T) {
	correct := "114"
	ans := part1(testInput, "g")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart1H(t *testing.T) {
	correct := "65412"
	ans := part1(testInput, "h")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart1I(t *testing.T) {
	correct := "65079"
	ans := part1(testInput, "i")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart1X(t *testing.T) {
	correct := "123"
	ans := part1(testInput, "x")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart1Y(t *testing.T) {
	correct := "456"
	ans := part1(testInput, "y")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart1B(t *testing.T) {
	correct := "96"
	ans := part1(testInput, "b")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

func TestPart1A(t *testing.T) {
	correct := "65439"
	ans := part1(testInput, "a")
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}
