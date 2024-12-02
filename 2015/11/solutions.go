package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"slices"

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
	// parsed := aoc.ParseInput(partInput)

	return "Answer 1"
}

func part2(partInput string) string {
	// parsed := aoc.ParseInput(partInput)

	return "Answer 2"
}

func incrementPassword(password string) string {
	// Incrementing relies on converting individual runes to an int (a=97, b=98... etc)
	runeStart := 'a'
	alphabetLengthMinusA := 25

	passwordRunes := []rune(password)
	slices.Reverse(passwordRunes)
	
	incrementing := true
	incrementingIndex := 0

	for incrementing {
		currentRune := passwordRunes[incrementingIndex]

		// Break out of the loop if there's no overflow - we're done incrementing
		if currentRune < runeStart + rune(alphabetLengthMinusA) {
			passwordRunes[incrementingIndex]++
			incrementing = false
			continue;
		}

		// If this is the last character we can't increment further
		if incrementingIndex == len(password) - 1 {
			incrementing = false
			continue
		}

		// Last case set the current rune to 'a' and move to the next letter
		passwordRunes[incrementingIndex] = runeStart
		incrementingIndex++
	}

	slices.Reverse(passwordRunes)

	return string(passwordRunes)
}

func passwordIsValid(password string) bool {
	runeBlacklist := []rune{'i', 'o', 'l'}
	
}
