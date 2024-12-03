package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"math"
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
	reports := strings.Split(parsed, "\n")

	safeReportCount := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReportCount++
		}
	}

	return strconv.Itoa(safeReportCount)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	reports := strings.Split(parsed, "\n")

	safeReportCount := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReportCount++
			continue
		}

		if isReportSafeWithOneRemoved(report) {
			safeReportCount++
		}
	}

	return strconv.Itoa(safeReportCount)
}

func isReportSafe(report string) bool {
	levels := strings.Split(report, " ")

	safe := true
	previousLevelDelta := 0

	for i, level := range levels {
		// Nothing to compare on the first level
		if i == 0 {
			continue
		}

		curr, _ := strconv.Atoi(level)
		prev, _ := strconv.Atoi(levels[i-1])

		// Get the difference between the current and previous level
		levelDelta := curr - prev
		absoluteDelta := math.Abs(float64(levelDelta))

		// Check of the delta is greater than 3 or zero
		if absoluteDelta > 3 || absoluteDelta == 0 {
			safe = false
			break
		}

		// Check if the signs on the delta are different (i.e a change from increasing -> decreasing)
		// math.Signbit will return `true` for negative values
		currentDeltaIsDecreasing := math.Signbit(float64(levelDelta))
		previousDeltaIsDecreasing := math.Signbit(float64(previousLevelDelta))

		// We don't have a previous delta until we hit at least item 3
		if i > 1 && currentDeltaIsDecreasing != previousDeltaIsDecreasing {
			safe = false
			break
		}

		previousLevelDelta = levelDelta
	}

	return safe
}

func isReportSafeWithOneRemoved(report string) bool {
	levels := strings.Split(report, " ")
	safe := false

	for i := range levels {
		levelsWithIndexRemoved := slices.Concat(levels[:i], levels[i+1:])
		newReport := strings.Join(levelsWithIndexRemoved, " ")

		if isReportSafe(newReport) {
			safe = true
			break
		}
	}

	return safe
}
