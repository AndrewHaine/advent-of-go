package main

import (
	"andrewhaine/advent-of-go/util/aoc"
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

func part1(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	rows := strings.Split(parsed, "\n")

	totalCodeLength, totalMemoryLength := 0, 0

	for _, row := range rows {
		totalCodeLength += getCodeCharLength(row)
		totalMemoryLength += getMemoryCharLength(row)
	}

	total := totalCodeLength - totalMemoryLength

	return strconv.Itoa(total)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	rows := strings.Split(parsed, "\n")

	totalCodeLength, totalMemoryLength := 0, 0

	for _, row := range rows {
		encodedRow := strconv.Quote(row)
		totalCodeLength += getCodeCharLength(encodedRow)
		totalMemoryLength += getMemoryCharLength(encodedRow)
	}

	total := totalCodeLength - totalMemoryLength

	return strconv.Itoa(total)
}

func getCodeCharLength(inputString string) int {
	return len(inputString)
}

func getMemoryCharLength(inputString string) int {
	unquoted, _ := strconv.Unquote(inputString)
	return len(unquoted)
}
