package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"reflect"
	"slices"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type Coordinate struct {
	x int
	y int
}

type WordSearch map[int][]string

// These are the movements/deltas on the x and y axis that are required to check all adjacent coordinates
var adjacentDeltas = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{-1, -1},
	{-1, 1},
	{1, 1},
	{1, -1},
}

var diagonalDeltaPairs = [][2][2]int{
	{{-1, -1}, {1, 1}},
	{{-1, 1}, {1, -1}},
}

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
	wordsearch := generateWordsearch(parsed)

	total := 0

	for y, row := range wordsearch {
		for x, value := range row {
			if value == "X" {
				coordinate := Coordinate{x: x, y: y}
				total += getXmasStringCountFromSource(coordinate, wordsearch)
			}
		}
	}

	return strconv.Itoa(total)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	wordsearch := generateWordsearch(parsed)

	total := 0

	for y, row := range wordsearch {
		for x, value := range row {
			if value == "A" {
				coordinate := Coordinate{x: x, y: y}
				isValid := isDiagonalXmas(coordinate, wordsearch)
				if isValid {
					total++
				}
			}
		}
	}

	return strconv.Itoa(total)
}

func generateWordsearch(stringInput string) WordSearch {
	wordsearch := WordSearch{}
	rows := strings.Split(stringInput, "\n")

	for i, row := range rows {
		wordsearch[i] = strings.Split(row, "")
	}

	return wordsearch
}

func isValidCoordinate(coordinate Coordinate, wordsearch WordSearch) bool {
	return (coordinate.y > -1 &&
		coordinate.x > -1 &&
		coordinate.y < len(wordsearch) &&
		coordinate.x < len(wordsearch[coordinate.y]))
}

func getXmasStringCountFromSource(coordinate Coordinate, wordsearch WordSearch) int {
	xmasString := []string{"X", "M", "A", "S"}
	count := 0

	// Loop through each adjacent letter
	for _, delta := range adjacentDeltas {

		// We already know we've hit the first letter
		i := 1

		// Loop through the slice of letters we need
		for i < len(xmasString) {

			// Check the next coordinate by multiplying the deltas by the current iteration
			newCoordinate := Coordinate{x: coordinate.x + (delta[0] * i), y: coordinate.y + (delta[1] * i)}

			// Escape hatch for when we've hit an edge
			if !isValidCoordinate(newCoordinate, wordsearch) {
				break
			}

			// Check the letter at the new position
			wordsearchValue := wordsearch[newCoordinate.y][newCoordinate.x]

			// Escape hatch for when we've it an invalid letter
			if wordsearchValue != xmasString[i] {
				break
			}

			// The letter we found is the one we're looking for - carry on
			i++
		}

		// If we iterated all the way through the letter, XMAS was found in this direction
		if i == len(xmasString) {
			count++
			continue
		}
	}

	return count
}

func isDiagonalXmas(coordinate Coordinate, wordsearch WordSearch) bool {
	valid := true
	requiredOrderedOuterLetters := []string{"M", "S"}

	for _, deltaPair := range diagonalDeltaPairs {
		letters := []string{}

		for _, delta := range deltaPair {
			newCoordinate := Coordinate{x: coordinate.x + delta[0], y: coordinate.y + delta[1]}

			if !isValidCoordinate(newCoordinate, wordsearch) {
				break
			}

			letter := wordsearch[newCoordinate.y][newCoordinate.x]
			letters = append(letters, letter)
		}

		slices.Sort(letters)

		if !reflect.DeepEqual(letters, requiredOrderedOuterLetters) {
			valid = false
			break
		}
	}

	return valid
}
