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

	pipeMap, startCoordinate := generatePipeMap(parsed)
	length, _ := findLoop(pipeMap, startCoordinate)

	furthest := length / 2

	return strconv.Itoa(furthest)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)

	pipeMap, startCoordinate := generatePipeMap(parsed)
	length, vertices := findLoop(pipeMap, startCoordinate)

	area := getShoelaceArea(vertices)

	// Rearrange pick's theorem to get the number of internal points
	inner := area - (length / 2) + 1

	return strconv.Itoa(inner)
}

func generatePipeMap(raw string) (pipeMap [][]rune, startCoordinate Coordinate) {
	rows := strings.Split(raw, "\n")

	for i, row := range rows {
		rowPipes := []rune{}
		for j, col := range row {
			rowPipes = append(rowPipes, col)
			if col == 'S' {
				startCoordinate = Coordinate{j, i}
			}
		}
		pipeMap = append(pipeMap, rowPipes)
	}

	return
}

func isValidCoordinate(coordinate Coordinate, pipeMap [][]rune) bool {
	if coordinate.x == -1 || coordinate.y == -1 {
		return false
	}

	if coordinate.y >= len(pipeMap) {
		return false
	}

	if coordinate.x >= len(pipeMap[coordinate.y]) {
		return false
	}

	return true
}

func flipDirection(direction [2]int) [2]int {
	for i, delta := range direction {
		if delta == -1 {
			direction[i] = 1
		} else if delta == 1 {
			direction[i] = -1
		}
	}

	return direction
}

func findLoop(pipeMap [][]rune, startCoordinate Coordinate) (length int, vertices []Coordinate) {
	length = 0
	currentPipe := pipeMap[startCoordinate.y][startCoordinate.x]
	currentCoordinate := startCoordinate
	visited := map[Coordinate]bool{}
	vertices = []Coordinate{}

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var possibleDirections = map[rune][][2]int{
		'|': {directions[0], directions[2]},
		'L': {directions[2], directions[1]},
		'J': {directions[2], directions[3]},
		'-': {directions[1], directions[3]},
		'F': {directions[0], directions[1]},
		'7': {directions[3], directions[0]},
	}

	for currentPipe != 'S' || length < 2 {
		if currentPipe == 'S' {
			for _, direction := range directions {
				newCoordinate := Coordinate{currentCoordinate.x + direction[0], currentCoordinate.y + direction[1]}

				if !isValidCoordinate(newCoordinate, pipeMap) {
					continue
				}

				newPipe := pipeMap[newCoordinate.y][newCoordinate.x]

				if slices.Contains(possibleDirections[newPipe], flipDirection(direction)) {
					visited[newCoordinate] = true
					currentCoordinate = newCoordinate
					currentPipe = newPipe
					break
				}
			}
		} else {
			possibleCurrentDirections := possibleDirections[currentPipe]

			for _, direction := range possibleCurrentDirections {
				newCoordinate := Coordinate{currentCoordinate.x + direction[0], currentCoordinate.y + direction[1]}

				if !isValidCoordinate(newCoordinate, pipeMap) {
					continue
				}

				if visited[newCoordinate] {
					continue
				}

				newPipe := pipeMap[newCoordinate.y][newCoordinate.x]

				if length < 2 && newPipe == 'S' {
					continue
				}

				visited[newCoordinate] = true
				currentCoordinate = newCoordinate
				currentPipe = newPipe
				break
			}
		}

		cornerPipes := []rune{'F', 'J', 'L', '7', 'S'}

		if slices.Contains(cornerPipes, currentPipe) {
			vertices = append(vertices, currentCoordinate)
		}

		length += 1
	}

	return
}

func getShoelaceArea(coordinates []Coordinate) (area int) {
	area = 0

	xy := 0
	yx := 0

	for i := 0; i < len(coordinates); i += 1 {
		if i == len(coordinates)-1 {
			xy += coordinates[i].x * coordinates[0].y
			yx += coordinates[i].y * coordinates[0].x
			continue
		}

		xy += coordinates[i].x * coordinates[i+1].y
		yx += coordinates[i].y * coordinates[i+1].x
	}

	area = int(math.Abs(float64(xy-yx)) / 2)

	return
}
