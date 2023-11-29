package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"andrewhaine/advent-of-go/util/sliceutil"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

var alphabet = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	all := strings.Split(parsed, "\n")

	sum := 0

	for _, rucksack := range all {
		compartmentSize := len(rucksack) / 2
		var compartment1, compartment2 = rucksack[:compartmentSize], rucksack[compartmentSize:]
		var itemType rune

		for _, letter := range compartment1 {
			if strings.Contains(compartment2, string(letter)) {
				itemType = letter
				break
			}
		}

		index := strings.Index(alphabet, string(itemType))
		sum += index
	}

	return strconv.Itoa(sum)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	all := strings.Split(parsed, "\n")

	sum := 0

	groups := sliceutil.Chunk[string](all, 3)

	for _, group := range groups {
		var itemType rune

		for _, letter := range group[0] {
			inTwo, inThree := strings.ContainsRune(group[1], letter), strings.ContainsRune(group[2], letter)

			if inTwo && inThree {
				itemType = letter
				break
			}
		}

		index := strings.Index(alphabet, string(itemType))
		sum += index
	}

	return strconv.Itoa(sum)
}
