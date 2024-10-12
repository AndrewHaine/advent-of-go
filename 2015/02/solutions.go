package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"regexp"
	"sort"
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
	presents := strings.Split(parsed, "\n")

	totalWrappingPaperAreaSqFeet := 0;

	for _, present := range presents {
		
		l, w, h := getPresentMeasurements(present)
		presentAreaSqFeet := 2*l*w + 2*w*h + 2*h*l

		// Find the smallest side and add the area
		measurements := []int{l, w, h}
		sort.Ints(measurements);

		presentAreaSqFeet += (measurements[0] * measurements[1])

		totalWrappingPaperAreaSqFeet += presentAreaSqFeet
	}

	return strconv.Itoa(totalWrappingPaperAreaSqFeet)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	presents := strings.Split(parsed, "\n")

	totalRibbonLengthFeet := 0

	for _, present := range presents {
		l, w, h := getPresentMeasurements(present)

		perimiters := []int{l + w, l + h, w + h}
		sort.Ints(perimiters)

		presentVolume := l*w*h
		wrapAroundRibbonLength := 2 * perimiters[0]

		totalRibbonLengthFeet += presentVolume + wrapAroundRibbonLength
	}

	return strconv.Itoa(totalRibbonLengthFeet)
}

func getPresentMeasurements(present string) (length int, width int, height int) {
	dimensionsRegex := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)
	matches := dimensionsRegex.FindStringSubmatch(present)
	
	l, _ := strconv.Atoi(matches[1])
	w, _ := strconv.Atoi(matches[2])
	h, _ := strconv.Atoi(matches[3])

	return l, w, h
}
