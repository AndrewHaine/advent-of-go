package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"andrewhaine/advent-of-go/util/mathutil"
	"fmt"
	"regexp"
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
	rows := strings.Split(aoc.ParseInput(partInput), "\n")
	count := 0

	for _, row := range rows {
		count += getCombinationsForRow(row)
	}

	return strconv.Itoa(count)
}

func part2(partInput string) string {
	rows := strings.Split(aoc.ParseInput(partInput), "\n")
	count := 0

	row := unfoldRow(rows[0])
	fmt.Println(row)
	count += getCombinationsForRow(row)

	// for _, row := range rows {
	// 	row = unfoldRow(row)
	// 	count += getCombinationsForRow(row)
	// }

	return strconv.Itoa(count)
}

func parseRow(row string) (arrangement string, damagedGroups []int) {
	split := strings.Split(row, " ")
	arrangement = split[0]

	damagedGroups = []int{}
	splitGroups := strings.Split(split[1], ",")
	for _, group := range splitGroups {
		groupSize, _ := strconv.Atoi(group)
		damagedGroups = append(damagedGroups, groupSize)
	}

	return
}

func unfoldRow(row string) string {
	split := strings.Split(row, " ")

	arrangement := split[0]
	damagedGroups := split[1]

	for i := 0; i < 4; i += 1 {
		arrangement += "?" + arrangement
		damagedGroups += "," + damagedGroups
	}

	return arrangement + " " + damagedGroups
}

func parseArrangement(arrangement string) (unknownPositions []int, knownCount int) {
	unknownPositions = []int{}
	knownCount = 0

	for i, char := range arrangement {
		if char == '?' {
			unknownPositions = append(unknownPositions, i)
		}

		if char == '#' {
			knownCount += 1
		}
	}

	return
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func getCombinationsForRow(row string) int {
	count := 0

	arrangement, damagedGroups := parseRow(row)
	unknownPositions, knownCount := parseArrangement(arrangement)
	totalDamagedSprings := mathutil.SumInts(damagedGroups)

	springsToAdd := totalDamagedSprings - knownCount

	permutations := binaryPermutations(len(unknownPositions), springsToAdd)

	for _, permutation := range permutations {
		parsedArrangement := arrangement
		for i, value := range permutation {
			if value == 1 {
				parsedArrangement = replaceAtIndex(parsedArrangement, '#', unknownPositions[i])
			} else {
				parsedArrangement = replaceAtIndex(parsedArrangement, '.', unknownPositions[i])
			}
		}
		matches := checkPatternMatch(parsedArrangement, damagedGroups)
		if matches {
			count += 1
		}
	}

	return count
}

func checkPatternMatch(arrangement string, damagedGroups []int) bool {
	regexString := `(\.+)?`
	for i, group := range damagedGroups {
		if i != 0 {
			regexString += `\.+`
		}
		regexString += `#{` + strconv.Itoa(group) + `}`
	}
	regexString += `(.+)?`

	checkRegex := regexp.MustCompile(regexString)

	return checkRegex.MatchString(arrangement)
}

func binaryPermutations(n, onCount int) [][]int {
	var result [][]int
	current := make([]int, n)
	generateHelperWithOnes(&result, current, 0, onCount)
	return result
}

func generateHelperWithOnes(result *[][]int, current []int, index, onCount int) {
	if index == len(current) {
		count := 0
		for _, num := range current {
			if num == 1 {
				count += 1
			}
		}
		if count == onCount {
			perm := make([]int, len(current))
			copy(perm, current)
			*result = append(*result, perm)
		}
		return
	}

	current[index] = 0
	generateHelperWithOnes(result, current, index+1, onCount)

	current[index] = 1
	generateHelperWithOnes(result, current, index+1, onCount)
}
