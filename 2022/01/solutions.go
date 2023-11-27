package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"flag"
	"slices"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Part 1 or 2")
	flag.Parse()

	if part == 1 {
		aoc.PrintSolution(part1(input))
	} else {
		aoc.PrintSolution(part2(input))
	}
}

func part1(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	elves := groupElves(parsed)

	return strconv.Itoa(slices.Max(elves))
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	elves := groupElves(parsed)

	slices.Sort(elves)

	largest3elves := elves[len(elves)-3:]

	totalCalories := 0
	for _, elf := range largest3elves {
		totalCalories = totalCalories + elf
	}

	return strconv.Itoa(totalCalories)
}

func groupElves(allElves string) []int {
	groups := strings.Split(allElves, "\n\n")

	elves := []int{}

	for _, group := range groups {
		calories := strings.Split(group, "\n")

		elfCalories := 0
		for _, calorie := range calories {
			calorieVal, _ := strconv.Atoi(calorie)
			elfCalories = elfCalories + calorieVal
		}

		elves = append(elves, elfCalories)
	}

	return elves
}
