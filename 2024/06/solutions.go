package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"slices"
	"sort"
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

type LabMap map[int][]string
type VisitedMap map[int][]string

type Delta [2]int
type DeltaWithLabel struct {
	delta Delta
	label string
}

var DELTA_UP = Delta{0, -1}
var DELTA_RIGHT = Delta{1, 0}
var DELTA_DOWN = Delta{0, 1}
var DELTA_LEFT = Delta{-1, 0}

var ORDERED_DELTAS = []DeltaWithLabel{
	{delta: DELTA_UP, label: "^"},
	{delta: DELTA_RIGHT, label: ">"},
	{delta: DELTA_DOWN, label: "v"},
	{delta: DELTA_LEFT, label: "<"},
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
	labMap := generateLabMap(parsed)

	var visitedMap = VisitedMap{}

	runPatrol(labMap, &visitedMap)

	totalVisitedCoordinates := calculateNumberOfVisitedCoordinates(visitedMap)

	return strconv.Itoa(totalVisitedCoordinates)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	labMap := generateLabMap(parsed)

	var visitedMap = VisitedMap{}

	loopObstacleCount := runPatrol(labMap, &visitedMap)

	return strconv.Itoa(loopObstacleCount)
}

func generateLabMap(stringInput string) LabMap {
	labMap := LabMap{}
	rows := strings.Split(stringInput, "\n")

	for i, row := range rows {
		labMap[i] = strings.Split(row, "")
	}

	return labMap
}

func isValidCoordinate(coordinate Coordinate, labMap LabMap) bool {
	return (coordinate.y > -1 &&
		coordinate.x > -1 &&
		coordinate.y < len(labMap) &&
		coordinate.x < len(labMap[coordinate.y]))
}

func getInitialCoordinate(labMap LabMap) Coordinate {
	activeCoordinate := Coordinate{x: 0, y: 0}

	found := false

	for y, row := range labMap {
		if found {
			break
		}

		for x, value := range row {
			activeCoordinate.x = x
			activeCoordinate.y = y

			if value == "^" {
				found = true
				break
			}
		}
	}

	return activeCoordinate
}

func trackVisitedCoordinateWithDirection(visitedMap *VisitedMap, visitedCoordinate Coordinate, direction string) {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(strconv.Itoa(visitedCoordinate.y))
	stringBuilder.WriteString("_")
	stringBuilder.WriteString(direction)
	value := stringBuilder.String()

	if _, ok := (*visitedMap)[visitedCoordinate.x]; !ok {
		(*visitedMap)[visitedCoordinate.x] = []string{value}
		return
	}

	newSlice := append((*visitedMap)[visitedCoordinate.x], value)
	sort.Strings(newSlice)

	(*visitedMap)[visitedCoordinate.x] = slices.Compact(newSlice)
}

func calculateNumberOfVisitedCoordinates(visitedMap VisitedMap) int {
	totalVisited := 0

	for _, visitedRows := range visitedMap {
		rowYCoordinates := []string{}

		for _, yCoordinate := range visitedRows {
			parts := strings.Split(yCoordinate, "_")
			rowYCoordinates = append(rowYCoordinates, parts[0])
		}

		uniqueVisitedRows := slices.Compact(rowYCoordinates)
		totalVisited += len(uniqueVisitedRows)
	}

	return totalVisited
}

func runPatrol(labMap LabMap, visitedMap *VisitedMap) int {
	initialDeltaIndex := 0
	loopObstacleCount := 0

	found := false
	activeDeltaIndex := initialDeltaIndex
	activeCoordinate := getInitialCoordinate(labMap)

	for !found {
		activeDelta := ORDERED_DELTAS[activeDeltaIndex]
		trackVisitedCoordinateWithDirection(visitedMap, activeCoordinate, activeDelta.label)

		next := Coordinate{x: activeCoordinate.x + activeDelta.delta[0], y: activeCoordinate.y + activeDelta.delta[1]}

		// This indicates we're about to leave the lab!
		if !isValidCoordinate(next, labMap) {
			found = true
			continue
		}

		// Loop through the next ordered deltas until we find a new path
		newActiveDeltaIndex := activeDeltaIndex + 1

		if newActiveDeltaIndex == len(ORDERED_DELTAS) {
			newActiveDeltaIndex = 0
		}

		// Check that the next value isn't an obstacle
		nextValue := labMap[next.y][next.x]
		if nextValue != "#" {
			activeCoordinate = next

			// The idea here is that we've tracked where we've visited alongside a direction, meaning
			// if we make a right turn and we've visited that coordinate in the proposed direction
			// then placing an obstacle in front of us yields a loop
			nextActiveDeltaWithObstacle := ORDERED_DELTAS[newActiveDeltaIndex]
			nextWithObstacle := Coordinate{
				x: activeCoordinate.x + nextActiveDeltaWithObstacle.delta[0],
				y: activeCoordinate.y + nextActiveDeltaWithObstacle.delta[1],
			}

			if !isValidCoordinate(nextWithObstacle, labMap) {
				continue
			}

			var stringBuilder strings.Builder

			stringBuilder.WriteString(strconv.Itoa(nextWithObstacle.y))
			stringBuilder.WriteString("_")
			stringBuilder.WriteString(nextActiveDeltaWithObstacle.label)
			value := stringBuilder.String()

			visitedX, ok := (*visitedMap)[nextWithObstacle.x]

			if ok && slices.Contains(visitedX, value) {
				loopObstacleCount++
			}

			continue
		}

		activeDeltaIndex = newActiveDeltaIndex
	}

	return loopObstacleCount
}
