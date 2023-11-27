package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"flag"

	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Part 1 or 2")
	flag.Parse()

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
