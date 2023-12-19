package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"math"
	"regexp"
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
	instructions := strings.Split(parsed, "\n")

	currentPosition := Coordinate{0, 0}
	trench := map[Coordinate]bool{}

	for _, instruction := range instructions {
		direction, distance, _ := parseInstruction(instruction)

		for i := distance; i > 0; i -= 1 {
			switch direction {
			case "R":
				currentPosition.x += 1
			case "L":
				currentPosition.x -= 1
			case "U":
				currentPosition.y -= 1
			case "D":
				currentPosition.y += 1
			}
			trench[currentPosition] = true
		}
	}

	floodFill(trench, Coordinate{1, 1})

	return strconv.Itoa(len(trench))
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	instructions := strings.Split(parsed, "\n")

	coordinates := []Coordinate{}
	currentPosition := Coordinate{0, 0}
	trenchArea := 0

	for _, instruction := range instructions {
		_, _, colour := parseInstruction(instruction)

		distance, _ := strconv.ParseInt(colour[:5], 16, 0)
		direction := colour[5]

		trenchArea += int(distance)

		switch direction {
		case '0':
			currentPosition.x += int(distance)
		case '2':
			currentPosition.x -= int(distance)
		case '3':
			currentPosition.y -= int(distance)
		case '1':
			currentPosition.y += int(distance)
		}
		coordinates = append(coordinates, currentPosition)
	}

	total := getShoelaceArea(coordinates) + trenchArea/2 + 1

	return strconv.Itoa(total)
}

func parseInstruction(instruction string) (direction string, distance int, colour string) {
	regex := regexp.MustCompile(`([RLUD])\s(\d+)\s\(#([a-f,0-9]+)\)`)
	results := regex.FindStringSubmatch(instruction)

	direction = results[1]
	distance, _ = strconv.Atoi(results[2])
	colour = results[3]

	return
}

func floodFill(trench map[Coordinate]bool, coordinate Coordinate) {
	trench[coordinate] = true

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for _, delta := range directions {
		nextCoordinate := Coordinate{coordinate.x + delta[0], coordinate.y + delta[1]}

		if _, ok := trench[nextCoordinate]; !ok {
			floodFill(trench, nextCoordinate)
		}
	}
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
