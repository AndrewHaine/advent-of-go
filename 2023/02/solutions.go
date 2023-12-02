package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"regexp"
	"slices"
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
	games := strings.Split(parsed, "\n")

	total := 0

	for i, game := range games {
		splitRow := strings.Split(game, ": ")
		colourCounts := getNumbersPerColour(splitRow[1])

		maxBlue := slices.Max(colourCounts["blue"])
		maxRed := slices.Max(colourCounts["red"])
		maxGreen := slices.Max(colourCounts["green"])

		if maxBlue <= 14 && maxGreen <= 13 && maxRed <= 12 {
			total += i + 1
		}
	}

	return strconv.Itoa(total)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	games := strings.Split(parsed, "\n")

	total := 0

	for _, game := range games {
		splitRow := strings.Split(game, ": ")
		colourCounts := getNumbersPerColour(splitRow[1])

		minBlue := slices.Max(colourCounts["blue"])
		minRed := slices.Max(colourCounts["red"])
		minGreen := slices.Max(colourCounts["green"])

		power := minBlue * minRed * minGreen

		total += power
	}

	return strconv.Itoa(total)
}

func getNumbersPerColour(gameOutcomes string) map[string][]int {
	cubeRegexp := regexp.MustCompile(`\d+\s(red|blue|green)`)
	colours := cubeRegexp.FindAllString(gameOutcomes, -1)

	colourNumbers := map[string][]int{}

	for _, colour := range colours {
		colourNumberPair := strings.Split(colour, " ")
		count, name := colourNumberPair[0], colourNumberPair[1]
		cubeCount, _ := strconv.Atoi(count)
		colourNumbers[name] = append(colourNumbers[name], cubeCount)
	}

	return colourNumbers
}
