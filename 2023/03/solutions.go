package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"strconv"
	"strings"
	"unicode"

	_ "embed"
)

//go:embed input.txt
var input string

type Coordinate struct {
	x int
	y int
}

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
	schematic := map[int][]rune{}

	for i, row := range rows {
		rowSlice := []rune(row)
		schematic[i] = rowSlice
	}

	total := 0

	for y, row := range schematic {
		x := 0
		for x < len(row) {
			c := Coordinate{x, y}
			if unicode.IsDigit(schematic[c.y][c.x]) {
				group, isPart := checkPart(c, schematic)
				if isPart {
					partValue, _ := strconv.Atoi(group)
					total += partValue
				}
				x += len(group)
			}
			x += 1
		}
	}

	return strconv.Itoa(total)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)

	rows := strings.Split(parsed, "\n")
	schematic := map[int][]rune{}

	for i, row := range rows {
		rowSlice := []rune(row)
		schematic[i] = rowSlice
	}

	total := 0

	for y, row := range schematic {
		x := 0
		for x < len(row) {
			c := Coordinate{x, y}
			if schematic[c.y][c.x] == '*' {
				ratio, isGear := checkGear(c, schematic)
				if isGear {
					total += ratio
				}
			}
			x += 1
		}
	}

	return strconv.Itoa(total)
}

// Returns the value of the number and whether there's a symbol around it
func checkPart(c Coordinate, schematic map[int][]rune) (string, bool) {
	part, _ := getPartFromCoordinate(c, schematic)
	partString := strconv.Itoa(part)
	partLength := len(partString)

	var neighbourContainsSymbol = false

	currentX := c.x

	for partLength > 0 {
		for _, delta := range adjacentDeltas {
			adjacentCoordinate := Coordinate{currentX + delta[0], c.y + delta[1]}

			if !isValidCoordinate(adjacentCoordinate, schematic) {
				continue
			}

			value := schematic[adjacentCoordinate.y][adjacentCoordinate.x]

			if !unicode.IsDigit(rune(value)) && value != '.' {
				neighbourContainsSymbol = true
				break
			}
		}
		currentX += 1
		partLength -= 1
	}

	return partString, neighbourContainsSymbol
}

func checkGear(c Coordinate, schematic map[int][]rune) (int, bool) {
	isGear := false
	ratio := 0

	neighbourCoordinates := []Coordinate{}
	partNeighbours := []int{}

	for _, delta := range adjacentDeltas {
		adjacentCoordinate := Coordinate{c.x + delta[0], c.y + delta[1]}

		if !isValidCoordinate(adjacentCoordinate, schematic) {
			continue
		}

		value := schematic[adjacentCoordinate.y][adjacentCoordinate.x]

		if unicode.IsDigit(rune(value)) {
			neighbourCoordinates = append(neighbourCoordinates, adjacentCoordinate)
		}
	}

	visitedCoordinates := map[Coordinate]bool{}

	for _, neighbour := range neighbourCoordinates {
		// If we've already seen this one don't process it again, this will result in parts registering multiple times
		if visitedCoordinates[neighbour] {
			continue
		}

		partValue, visitedDuringCheck := getPartFromCoordinate(neighbour, schematic)

		for _, visited := range visitedDuringCheck {
			visitedCoordinates[visited] = true
		}

		partNeighbours = append(partNeighbours, partValue)
	}

	if len(partNeighbours) == 2 {
		isGear = true
		ratio = partNeighbours[0] * partNeighbours[1]
	}

	return ratio, isGear
}

func isValidCoordinate(coordinate Coordinate, schematic map[int][]rune) bool {
	return (coordinate.y > -1 &&
		coordinate.x > -1 &&
		coordinate.y < len(schematic) &&
		coordinate.x < len(schematic[coordinate.y]))
}

func getPartFromCoordinate(coordinate Coordinate, schematic map[int][]rune) (int, []Coordinate) {
	visitedCoordinates := []Coordinate{}
	startX := coordinate.x
	lowestX, highestX := startX, startX

	for lowestX != 0 && unicode.IsDigit(schematic[coordinate.y][lowestX-1]) {
		lowestX -= 1
	}

	for highestX+1 < len(schematic[0]) && unicode.IsDigit(schematic[coordinate.y][highestX+1]) {
		highestX += 1
	}

	partDigits := []rune{}

	for lowestX <= highestX {
		checkingCoordinate := Coordinate{lowestX, coordinate.y}
		visitedCoordinates = append(visitedCoordinates, checkingCoordinate)
		partDigits = append(partDigits, schematic[checkingCoordinate.y][checkingCoordinate.x])
		lowestX += 1
	}

	partString := string(partDigits)
	partValue, _ := strconv.Atoi(partString)
	return partValue, visitedCoordinates
}
