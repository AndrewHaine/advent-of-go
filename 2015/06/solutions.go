package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Bulbs [1_000_000]int

type Coordinate struct {
	x int
	y int
}

type BulbCallbackFunc func (bulbs *Bulbs, bulbIndex int, action string)

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

	litBulbCount := 0

	var bulbs Bulbs

	for _, instruction := range instructions {
		processInstruction(instruction, &bulbs, performActionOnBulbPart1)
	}

	for _, bulb := range bulbs {
		if bulb == 1 {
			litBulbCount++
		}
	}

	return strconv.Itoa(litBulbCount)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	instructions := strings.Split(parsed, "\n")

	bulbBrightness := 0

	var bulbs Bulbs

	for _, instruction := range instructions {
		processInstruction(instruction, &bulbs, performActionOnBulbPart2)
	}

	for _, bulb := range bulbs {
		bulbBrightness += bulb
	}

	return strconv.Itoa(bulbBrightness)
}

func processInstruction(instruction string, bulbs *Bulbs, performActionOnBulb BulbCallbackFunc) {
	rowLength := 1_000
	
	instructionRegex := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	parsedInstruction := instructionRegex.FindStringSubmatch(instruction)

	action := parsedInstruction[1]
	startX, _ := strconv.Atoi(parsedInstruction[2])
	startY, _ := strconv.Atoi(parsedInstruction[3])
	endX, _ := strconv.Atoi(parsedInstruction[4])
	endY, _ := strconv.Atoi(parsedInstruction[5])

	for y := startY; y <= endY; y++ {
		for x:= startX; x <= endX; x++ {
			i := getBulbFromCoordinate(Coordinate{x, y}, rowLength)
			performActionOnBulb(bulbs, i, action)
		}
	}
}

func performActionOnBulbPart1(bulbs *Bulbs, bulbIndex int, action string) {
	switch action {
	case "turn on":
		bulbs[bulbIndex] = 1;
	case "turn off":
		bulbs[bulbIndex] = 0
	case "toggle":
		if bulbs[bulbIndex] == 1 {
			bulbs[bulbIndex] = 0
		} else {
			bulbs[bulbIndex] = 1
		}
	}
}
func performActionOnBulbPart2(bulbs *Bulbs, bulbIndex int, action string) {
	switch action {
	case "turn on":
		bulbs[bulbIndex]++;
	case "turn off":
		bulbs[bulbIndex] = max(bulbs[bulbIndex] -1, 0)
	case "toggle":
		bulbs[bulbIndex] += 2;
	}
}

func getBulbFromCoordinate(coordinate Coordinate, rowLength int) int {
	// Getting the position of a coordinate is as simple as (rowLength * y) + x
	return (rowLength * coordinate.y) + coordinate.x
}
