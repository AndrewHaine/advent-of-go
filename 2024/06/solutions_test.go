package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var testInput string

/*
The map shows the current position of the guard with ^ (to indicate the guard is currently facing up from the perspective of the map). Any obstructions - crates, desks, alchemical reactors, etc. - are shown as #.

Lab guards in 1518 follow a very strict patrol protocol which involves repeatedly following these steps:

If there is something directly in front of you, turn right 90 degrees.
Otherwise, take a step forward.

This process continues for a while, but the guard eventually leaves the mapped area (after walking past a tank of universal solvent)

In this example, the guard will visit 41 distinct positions on your map.

Predict the path of the guard. How many distinct positions will the guard visit before leaving the mapped area?
*/
func TestPart1(t *testing.T) {
	correct := "41"
	ans := part1(testInput)
	if ans != correct {
		t.Errorf("part1() got %s; want %s", ans, correct)
	}
}

/*
Fortunately, they are pretty sure that adding a single new obstruction won't cause a time paradox. They'd like to place the new obstruction in such a way that the guard will get stuck in a loop, making the rest of the lab safe to search.

To have the lowest chance of creating a time paradox, The Historians would like to know all of the possible positions for such an obstruction. The new obstruction can't be placed at the guard's starting position - the guard is there right now and would notice.
*/
func TestPart2(t *testing.T) {
	correct := "6"
	ans := part2(testInput)
	if ans != correct {
		t.Errorf("part2() got %s; want %s", ans, correct)
	}
}
