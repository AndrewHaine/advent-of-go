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
	lists := generateLocationListsFromInput(parsed)

	for _, list := range lists {
		slices.Sort(list)
	}

	distance := 0

	for i, firstLocationId := range lists[0] {
		secondLocationId := lists[1][i]
		
		// Go std has no absolute calculation so we ensure we're getting a positive integer for the distance
		if secondLocationId > firstLocationId {
			distance += secondLocationId - firstLocationId
			continue
		}

		distance += firstLocationId - secondLocationId
	}

	return strconv.Itoa(distance)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	lists := generateLocationListsFromInput(parsed)

	frequencyMap := generateLocationFrequencyMap(lists[1])

	similarityScore := 0

	for _, locationId := range lists[0] {
		if frequency, ok := frequencyMap[locationId]; ok {
			similarityScore += locationId * frequency
		}
	}

	return strconv.Itoa(similarityScore)
}

func generateLocationListsFromInput(inputString string) [2][]int {
	locationRows := strings.Split(inputString, "\n")

	lists := [2][]int{{}, {}}

	for _, row := range locationRows {
		rowRegex := regexp.MustCompile(`^(\d+)\s+(\d+)$`)
		matches := rowRegex.FindStringSubmatch(row)

		first, _ := strconv.Atoi(matches[1])
		second, _ := strconv.Atoi(matches[2])

		lists[0] = append(lists[0], first)
		lists[1] = append(lists[1], second)
	}

	return lists
}

func generateLocationFrequencyMap(list []int) map[int]int {
	frequencyMap := map[int]int{}

	for _, locationId := range list {
		currentFrequency, ok := frequencyMap[locationId]

		if ok {
			frequencyMap[locationId] = currentFrequency + 1
		} else {
			frequencyMap[locationId] = 1
		}
	}

	return frequencyMap
}
