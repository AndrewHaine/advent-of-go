package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"regexp"
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
	timeDistanceMap := generateTimeDistanceMapFromInput(parsed)

	counts := []int{}

	for time, distance := range timeDistanceMap {
		winningTimes := 0
		timeToTest := 1

		for timeToTest < time {
			distanceForTime := (time - timeToTest) * timeToTest

			if distanceForTime > distance {
				winningTimes += 1
			}

			timeToTest += 1
		}

		counts = append(counts, winningTimes)
	}

	marginOfError := 1

	for _, count := range counts {
		marginOfError = marginOfError * count
	}

	return strconv.Itoa(marginOfError)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	rows := strings.Split(parsed, "\n")

	time := 0
	distance := 0

	for i, row := range rows {
		spacesRemoved := strings.ReplaceAll(row, " ", "")
		split := strings.Split(spacesRemoved, ":")
		value, _ := strconv.Atoi(split[1])

		if i == 0 {
			time = value
		}

		if i == 1 {
			distance = value
		}
	}

	timeToTest := 1
	winningTimes := 0

	for timeToTest < time {
		distanceForTime := (time - timeToTest) * timeToTest

		if distanceForTime > distance {
			winningTimes += 1
		}

		timeToTest += 1
	}

	return strconv.Itoa(winningTimes)
}

func generateTimeDistanceMapFromInput(input string) map[int]int {
	rows := strings.Split(input, "\n")

	times := []int{}
	distances := []int{}

	for i, row := range rows {
		rowRegex := regexp.MustCompile(`\d+`)
		games := rowRegex.FindAllStringSubmatch(row, -1)
		for _, game := range games {
			value, _ := strconv.Atoi(game[0])
			if i == 0 {
				times = append(times, value)
			}

			if i == 1 {
				distances = append(distances, value)
			}
		}
	}

	timeDistanceMap := map[int]int{}

	for i, time := range times {
		timeDistanceMap[time] = distances[i]
	}

	return timeDistanceMap
}
