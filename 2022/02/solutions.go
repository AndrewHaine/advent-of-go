package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"andrewhaine/advent-of-go/util/maputil"
	"strconv"
	"strings"

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

const (
	WIN  = 6
	DRAW = 3
	LOSS = 0
)

var handValues = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var winningCombinations = map[string]string{
	"X": "C",
	"Y": "A",
	"Z": "B",
}

var drawingCombinations = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var losingCombinations = map[string]string{
	"X": "B",
	"Y": "C",
	"Z": "A",
}

func part1(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	lines := strings.Split(parsed, "\n")

	score := 0

	for _, line := range lines {
		hands := strings.Split(line, " ")
		them, us := hands[0], hands[1]

		score += handValues[us]

		if winningCombinations[us] == them {
			score += WIN
		} else if drawingCombinations[us] == them {
			score += DRAW
		}
	}

	return strconv.Itoa(score)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	lines := strings.Split(parsed, "\n")

	invertedWinning := maputil.Invert(winningCombinations)
	invertedLosing := maputil.Invert(losingCombinations)
	invertedDrawing := maputil.Invert(drawingCombinations)

	score := 0

	for _, line := range lines {
		hands := strings.Split(line, " ")
		them, us := hands[0], hands[1]

		if us == "Y" {
			score += DRAW
			score += handValues[invertedDrawing[them]]
		} else if us == "Z" {
			score += WIN
			score += handValues[invertedWinning[them]]
		} else {
			score += handValues[invertedLosing[them]]
		}
	}

	return strconv.Itoa(score)
}
