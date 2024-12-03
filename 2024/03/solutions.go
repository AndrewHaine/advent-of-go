package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"regexp"
	"strconv"

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

	instructionsRegex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := instructionsRegex.FindAllString(parsed, -1)

	answer := 0

	for _, instruction := range matches {
		answer += getMulInstructionValue(instruction)
	}

	return strconv.Itoa(answer)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)

	instructionsRegex := regexp.MustCompile(`do\(\)|mul\(\d+,\d+\)|don't\(\)`)
	matches := instructionsRegex.FindAllString(parsed, -1)

	answer := 0
	counting := true

	for _, instruction := range matches {
		if instruction == "don't()" {
			counting = false
			continue
		}

		if instruction == "do()" {
			counting = true
			continue
		}

		if counting {
			answer += getMulInstructionValue(instruction)
		}
	}

	return strconv.Itoa(answer)
}

func getMulInstructionValue(instruction string) int {
	instructionRegex := regexp.MustCompile(`^mul\((\d+),(\d+)\)$`)
	instructionMatches := instructionRegex.FindStringSubmatch(instruction)

	first, _ := strconv.Atoi(instructionMatches[1])
	second, _ := strconv.Atoi(instructionMatches[2])

	return first * second
}
