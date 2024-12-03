package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"strconv"

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
	instructions := aoc.ParseInput(partInput)

	floor := 0

	for _, instruction := range instructions {
		if instruction == '(' {
			floor++
		} else {
			floor--
		}
	}

	return strconv.Itoa(floor)
}

func part2(partInput string) string {
	instructions := aoc.ParseInput(partInput)

	var basementPos int
	floor := 0

	for i, instruction := range instructions {
		if floor == -1 {
			basementPos = i
			break
		}

		if instruction == '(' {
			floor++
		} else {
			floor--
		}
	}

	return strconv.Itoa(basementPos)
}
