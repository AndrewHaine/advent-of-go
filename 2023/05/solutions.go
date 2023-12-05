package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"andrewhaine/advent-of-go/util/maputil"
	"andrewhaine/advent-of-go/util/sliceutil"
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

type RangeMapping struct {
	lower    int
	upper    int
	modifier int
}

func part1(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	categories := strings.Split(parsed, "\n\n")

	seedString := strings.Split(categories[0], ":")
	seeds := stringToNumberList(seedString[1])

	seedLocations := map[int]int{}

	for _, seed := range seeds {
		value := seed

		for _, mapping := range categories[1:] {
			value = getMappedValue(mapping, value)
		}

		seedLocations[seed] = value
	}

	locationValues := maputil.Values(seedLocations)
	lowestLocation := slices.Min(locationValues)

	return strconv.Itoa(lowestLocation)
}

/*
 Brute force all seeds (takes over 2 hours)
*/
// func part2(partInput string) string {
// 	parsed := aoc.ParseInput(partInput)
// 	categories := strings.Split(parsed, "\n\n")

// 	seedString := strings.Split(categories[0], ":")
// 	seeds := stringToNumberList(seedString[1])
// 	seedPairs := sliceutil.Chunk[int](seeds, 2)

// 	seedLocations := map[int]int{}

// 	for _, seedPair := range seedPairs {
// 		seedsInPair := []int{}

// 		for rangeIndex := 1; rangeIndex <= seedPair[1]; rangeIndex += 1 {
// 			seedsInPair = append(seedsInPair, seedPair[0]+rangeIndex)
// 		}

// 		for _, seed := range seedsInPair {
// 			value := seed

// 			for _, mapping := range categories[1:] {
// 				value = getMappedValue(mapping, value)
// 			}

// 			seedLocations[seed] = value
// 		}
// 	}

// 	locationValues := maputil.Values(seedLocations)
// 	lowestLocation := slices.Min(locationValues)

// 	return strconv.Itoa(lowestLocation)
// }

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	categories := strings.Split(parsed, "\n\n")

	seedString := strings.Split(categories[0], ":")
	seedPairs := sliceutil.Chunk(stringToNumberList(seedString[1]), 2)

	// Invert mappings, find a seed that matches the location
	matchingLocation := 0
	currentLocation := 0

	mappingBlocks := categories[1:]
	slices.Reverse(mappingBlocks)

	for matchingLocation < 1 {
		locationSeed := getSeedFromLocation(mappingBlocks, currentLocation)
		if isValidSeed(locationSeed, seedPairs) {
			matchingLocation = currentLocation
			break
		}
		currentLocation += 1
	}

	return strconv.Itoa(matchingLocation)
}

func stringToNumberList(numbersString string) []int {
	trimmed := strings.Trim(numbersString, " ")
	numberStrings := strings.Split(trimmed, " ")
	numbers := []int{}
	for _, numberString := range numberStrings {
		number, _ := strconv.Atoi(numberString)
		numbers = append(numbers, number)
	}
	return numbers
}

func getMappedValue(block string, value int) int {
	newValue := value
	rows := strings.Split(block, "\n")

	for _, mappingString := range rows[1:] {
		mapping := stringToNumberList(mappingString)
		destination, source, count := mapping[0], mapping[1], mapping[2]

		maxValue := source + count

		if newValue >= source && newValue <= maxValue {
			diff := newValue - source
			newValue = destination + diff
			break
		}
	}

	return newValue
}

func getSeedFromLocation(mappings []string, location int) int {
	value := location
	for _, mapping := range mappings {
		value = getInversMappedValue(mapping, value)
	}
	return value
}

func getInversMappedValue(block string, value int) int {
	newValue := value
	rows := strings.Split(block, "\n")

	for _, mappingString := range rows[1:] {
		mapping := stringToNumberList(mappingString)
		destination, source, count := mapping[0], mapping[1], mapping[2]

		if newValue <= destination+count && newValue >= destination {
			diff := newValue - destination
			newValue = source + diff
			break
		}
	}

	return newValue
}

func isValidSeed(seed int, seedPairs [][]int) bool {
	isValid := false

	for _, seedPair := range seedPairs {
		if seed >= seedPair[0] && seed <= seedPair[0]+seedPair[1] {
			isValid = true
			break
		}
	}

	return isValid
}
