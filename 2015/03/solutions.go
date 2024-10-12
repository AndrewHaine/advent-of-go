package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"slices"
	"sort"
	"strconv"

	_ "embed"
)

//go:embed input.txt
var input string

type Coordinate struct {
	x int
	y int
}

type VisitedMap map[int][]int

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

	visitedMap := VisitedMap{0: {0}}
	santaPosition := Coordinate{0, 0}

	for _, instruction := range parsed {
		santaPosition = moveCoordinate(santaPosition, instruction)
		trackVisitedCoordinate(&visitedMap, santaPosition)
	}

	totalVisitedHouses := calculateNumberOfVisitedHouses(visitedMap)

	return strconv.Itoa(totalVisitedHouses)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)

	visitedMap := VisitedMap{0: {0}}
	santaPosition := Coordinate{0, 0}
	roboSantaPosition := Coordinate{0, 0}

	for i, instruction := range parsed {
		if (i % 2 == 0) {
			santaPosition = moveCoordinate(santaPosition, instruction)
			trackVisitedCoordinate(&visitedMap, santaPosition)
			continue
		}

		roboSantaPosition = moveCoordinate(roboSantaPosition, instruction)
		trackVisitedCoordinate(&visitedMap, roboSantaPosition)
	}

	totalVisitedHouses := calculateNumberOfVisitedHouses(visitedMap)

	return strconv.Itoa(totalVisitedHouses)
}

func moveCoordinate(coordinate Coordinate, instruction rune) Coordinate {
	switch instruction {
	case '^':
		coordinate.y--
	case 'v':
		coordinate.y++
	case '>':
		coordinate.x++
	case '<':
		coordinate.x--
	}

	return coordinate
}

func trackVisitedCoordinate(visitedMap *VisitedMap, visitedCoordinate Coordinate) {
	if _, ok := (*visitedMap)[visitedCoordinate.x]; !ok {
		(*visitedMap)[visitedCoordinate.x] = []int{visitedCoordinate.y}
		return
	}

	newSlice := append((*visitedMap)[visitedCoordinate.x], visitedCoordinate.y)
	sort.Ints(newSlice)

	(*visitedMap)[visitedCoordinate.x] = slices.Compact(newSlice)
}

func calculateNumberOfVisitedHouses(visitedMap VisitedMap) int {
	totalVisitedHouses := 0

	for _, visitedRows := range visitedMap {
		totalVisitedHouses += len(visitedRows)
	}

	return totalVisitedHouses
}
