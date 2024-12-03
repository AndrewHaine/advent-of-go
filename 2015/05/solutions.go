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
	niceStringCount := 0
	unsortedStrings := strings.Split(aoc.ParseInput(partInput), "\n")

	for _, unsortedString := range unsortedStrings {
		if stringIsNicePart1(unsortedString) {
			niceStringCount++
		}
	}

	return strconv.Itoa(niceStringCount)
}

func part2(partInput string) string {
	niceStringCount := 0
	unsortedStrings := strings.Split(aoc.ParseInput(partInput), "\n")

	for _, unsortedString := range unsortedStrings {
		if stringIsNicePart2(unsortedString) {
			niceStringCount++
		}
	}

	return strconv.Itoa(niceStringCount)
}

func stringIsNicePart1(inputString string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	vowelCount := 0
	containsLetterPair := false
	disallowedStringPairs := [][2]rune{{'a', 'b'}, {'c', 'd'}, {'p', 'q'}, {'x', 'y'}}
	containsDisallowedStringPair := false

	for i, letter := range inputString {
		if slices.Contains(vowels, letter) {
			vowelCount++
		}

		// Checks from this point look back, not possible if we're at the head
		if i == 0 {
			continue
		}

		previousLetter := []rune(inputString)[i-1]

		if previousLetter == letter {
			containsLetterPair = true
		}

		for _, pair := range disallowedStringPairs {
			if pair[1] == letter && pair[0] == previousLetter {
				containsDisallowedStringPair = true
				break
			}
		}
	}

	return vowelCount >= 3 && containsLetterPair && !containsDisallowedStringPair
}

func stringIsNicePart2(inputString string) bool {
	containsRepeatingLetterPair := false
	containsSeparatedRepeatedLetter := false

	for i, letter := range inputString {
		if containsRepeatingLetterPair && containsSeparatedRepeatedLetter {
			break
		}

		if i < 2 {
			continue
		}

		if !containsSeparatedRepeatedLetter {
			secondPreviousLetter := []rune(inputString)[i-2]

			if secondPreviousLetter == letter {
				containsSeparatedRepeatedLetter = true
			}
		}

		if containsRepeatingLetterPair {
			continue
		}

		previousLetter := []rune(inputString)[i-1]
		pair := string([]rune{previousLetter, letter})

		lookBack := inputString[:i-1]

		if strings.Contains(lookBack, pair) {
			containsRepeatingLetterPair = true
		}
	}

	return containsRepeatingLetterPair && containsSeparatedRepeatedLetter
}
