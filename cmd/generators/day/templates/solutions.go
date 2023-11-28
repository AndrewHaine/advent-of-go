package main

import (
	"andrewhaine/advent-of-go/util/aoc"

	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	part := aoc.PartFlag()

	if part == 1 {
		aoc.PrintSolution(part1(input))
	} else {
		aoc.PrintSolution(part2(input))
	}
}

func part1(partInput string) string {
	// parsed := aoc.ParseInput(partInput)

	return "Answer 1"
}

func part2(partInput string) string {
	// parsed := aoc.ParseInput(partInput)

	return "Answer 2"
}
