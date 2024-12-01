package main

import (
	_ "embed"
	"testing"
)

func TestLookAndSayString1(t *testing.T) {
	correct := "11"
	ans := lookAndSayString("1")
	if ans != correct {
		t.Errorf("lookAndSayString(\"11\") got %s; want %s", ans, correct)
	}
}
func TestLookAndSayString2(t *testing.T) {
	correct := "21"
	ans := lookAndSayString("11")
	if ans != correct {
		t.Errorf("lookAndSayString(\"11\") got %s; want %s", ans, correct)
	}
}

func TestLookAndSayString3(t *testing.T) {
	correct := "1211"
	ans := lookAndSayString("21")
	if ans != correct {
		t.Errorf("lookAndSayString(\"21\") got %s; want %s", ans, correct)
	}
}

func TestLookAndSayString4(t *testing.T) {
	correct := "111221"
	ans := lookAndSayString("1211")
	if ans != correct {
		t.Errorf("lookAndSayString(\"1211\") got %s; want %s", ans, correct)
	}
}

func TestLookAndSayString5(t *testing.T) {
	correct := "312211"
	ans := lookAndSayString("111221")
	if ans != correct {
		t.Errorf("lookAndSayString(\"312211\") got %s; want %s", ans, correct)
	}
}
