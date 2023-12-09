package main

import (
	"andrewhaine/advent-of-go/util/aoc"
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

func part1(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	rows := strings.Split(parsed, "\n")

	sum := 0

	for _, row := range rows {
		rowList := stringToNumberList(row)
		sum += getNextNumberInSequence(rowList)
	}

	return strconv.Itoa(sum)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	rows := strings.Split(parsed, "\n")

	sum := 0

	for _, row := range rows {
		rowList := stringToNumberList(row)
		slices.Reverse(rowList)
		sum += getNextNumberInSequence(rowList)
	}

	return strconv.Itoa(sum)
}

func getNextNumberInSequence(row []int) int {
	rowDiffs := [][]int{row}
	allZero := false

	for !allZero {
		rowDiff := []int{}
		currentRow := rowDiffs[len(rowDiffs)-1]
		for i := 0; i < len(currentRow)-1; i += 1 {
			rowDiff = append(rowDiff, currentRow[i+1]-currentRow[i])
		}
		allZero = sliceutil.Every(rowDiff, func(index int, item int) bool {
			return item == 0
		})
		rowDiffs = append(rowDiffs, rowDiff)
	}

	slices.Reverse(rowDiffs)

	for i, rowDiff := range rowDiffs {
		if i == 0 {
			rowDiff = append(rowDiff, 0)
		} else {
			prevLast := rowDiffs[i-1][len(rowDiffs[i-1])-1]
			rowDiff = append(rowDiff, rowDiff[len(rowDiff)-1]+prevLast)
		}

		rowDiffs[i] = rowDiff
	}

	slices.Reverse(rowDiffs)

	return rowDiffs[0][len(rowDiffs[0])-1]
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
