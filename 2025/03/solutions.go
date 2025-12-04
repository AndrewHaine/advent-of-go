package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"slices"
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
	banks := parseBanks(parsed)

	total := 0

	for _, bank := range banks {
		total += getHighestJoltageSlidingWindow(bank, 2)
	}

	return strconv.Itoa(total)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	banks := parseBanks(parsed)

	total := 0

	for _, bank := range banks {
		total += getHighestJoltageSlidingWindow(bank, 12)
	}

	return strconv.Itoa(total)
}

func parseBanks(banksRaw string) [][]string {
	bankStrs := strings.Split(banksRaw, "\n")
	banks := [][]string{}

	for _, bankStr := range bankStrs {
		banks = append(banks, strings.Split(bankStr, ""))
	}

	return banks
}

func getHighestJoltageSlidingWindow(bank []string, windowSize int) int {
	window := bank[:windowSize]

	for i := range bank {
		if i+windowSize > len(bank) {
			break
		}

		window = evaluateBankWindow(window, bank[i:windowSize+i])
	}

	joltageString := ""

	for _, battery := range window {
		joltageString += battery
	}

	joltage, _ := strconv.Atoi(joltageString)
	return joltage
}

func evaluateBankWindow(window []string, newWindow []string) []string {
	for i := range window {
		if newWindow[i] > window[i] {
			window = slices.Concat(window[:i], newWindow[i:])
			break
		}
	}

	return window
}
