package main

import (
	"andrewhaine/advent-of-go/util/aoc"
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
	ranges, available := parseInput(parsed)

	tot := 0

	for _, availableId := range available {
		for _, rangeToCheck := range ranges {
			if availableId >= rangeToCheck[0] && availableId <= rangeToCheck[1] {
				tot++
				break
			}
		}
	}

	return strconv.Itoa(tot)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	ranges, _ := parseInput(parsed)

	return strconv.Itoa(part2Properly(ranges))
}

func parseInput(input string) (ranges [][2]int, available []int) {
	parts := strings.Split(input, "\n\n")

	for _, rangeRaw := range strings.Split(parts[0], "\n") {
		rangeParts := strings.Split(rangeRaw, "-")
		low, _ := strconv.Atoi(rangeParts[0])
		high, _ := strconv.Atoi(rangeParts[1])

		ranges = append(ranges, [2]int{low, high})
	}

	for _, availableRaw := range strings.Split(parts[1], "\n") {
		availableVal, _ := strconv.Atoi(availableRaw)
		available = append(available, availableVal)
	}

	return ranges, available
}

func part2BruteForce(ranges [][2]int) int {
	vals := []int{}

	for _, rangeToCheck := range ranges {
		curr := rangeToCheck[0]

		for curr <= rangeToCheck[1] {
			vals = append(vals, curr)
			curr++
		}
	}

	return len(vals)
}

func part2Properly(ranges [][2]int) int {
	// Sort the ranges by their low value
	slices.SortFunc(ranges, rangeSortFunc)

	tot := 0

	currRange := ranges[0]

	for _, rangeVal := range ranges[1:] {

		// If the next range starts after the current one ends move on to the next group and add currRange size to the total
		if rangeVal[0] > currRange[1] {
			tot += (currRange[1] - currRange[0]) + 1

			currRange = rangeVal
			continue
		}

		// If the next range ends after the current one, merge them together
		if rangeVal[1] > currRange[1] {
			currRange[1] = rangeVal[1]
			continue
		}
	}

	// Make sure the last range is accounted for
	tot += (currRange[1] - currRange[0]) + 1

	return tot
}

func rangeSortFunc(a [2]int, b [2]int) int {
	if a[0] > b[0] {
		return 1
	}

	return -1
}
