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
			coord := Coordinate{x, y}
			if unicode.IsDigit(schematic[coord.y][coord.x]) {
				group, isPart := checkPart(coord, schematic)
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
			coord := Coordinate{x, y}
			if schematic[coord.y][coord.x] == '*' {
				ratio, isGear := checkGear(coord, schematic)
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
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}

	groupRunes := []rune{}
	groupLength := 0

	for c.x+groupLength < len(schematic[c.y]) && unicode.IsDigit(schematic[c.y][c.x+groupLength]) {
		groupRunes = append(groupRunes, schematic[c.y][c.x+groupLength])
		groupLength += 1
	}

	var neighbourContainsSymbol = false

	currentX := c.x
	for groupLength > 0 {
		for _, delta := range directions {
			newX := currentX + delta[0]
			newY := c.y + delta[1]

			if newY == -1 || newX == -1 || newY == len(schematic) || newX == len(schematic[c.y]) {
				continue
			}

			value := schematic[newY][newX]

			if !unicode.IsDigit(rune(value)) && value != '.' {
				neighbourContainsSymbol = true
				break
			}
		}
		currentX += 1
		groupLength -= 1
	}

	groupString := string(groupRunes)

	return groupString, neighbourContainsSymbol
}

func checkGear(c Coordinate, schematic map[int][]rune) (int, bool) {
	isGear := false
	ratio := 0

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}

	coordsToCheck := []Coordinate{}
	partNeighbours := []int{}

	for _, delta := range directions {
		newX := c.x + delta[0]
		newY := c.y + delta[1]

		if newY == -1 || newX == -1 || newY == len(schematic) || newX == len(schematic[c.y]) {
			continue
		}

		value := schematic[newY][newX]

		if unicode.IsDigit(rune(value)) {
			coordsToCheck = append(coordsToCheck, Coordinate{newX, newY})
		}
	}

	seenCoords := map[Coordinate]bool{}

	for _, coord := range coordsToCheck {
		if seenCoords[coord] {
			continue
		}

		seenCoords[coord] = true

		// Look ahead and back to group numbers
		lowestX := coord.x
		highestX := coord.x

		for lowestX != 0 && unicode.IsDigit(schematic[coord.y][lowestX-1]) {
			lowestX -= 1
		}

		for highestX+1 < len(schematic[0]) && unicode.IsDigit(schematic[coord.y][highestX+1]) {
			highestX += 1
		}

		partDigits := []rune{}

		for lowestX <= highestX {
			checkingCoord := Coordinate{lowestX, coord.y}
			partDigits = append(partDigits, schematic[checkingCoord.y][checkingCoord.x])
			seenCoords[checkingCoord] = true
			lowestX += 1
		}

		partString := string(partDigits)
		partValue, _ := strconv.Atoi(partString)
		partNeighbours = append(partNeighbours, partValue)
	}

	if len(partNeighbours) == 2 {
		isGear = true
		ratio = partNeighbours[0] * partNeighbours[1]
	}

	return ratio, isGear
}
