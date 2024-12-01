package main

import (
	"andrewhaine/advent-of-go/util/aoc"
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
	inputString := aoc.ParseInput(partInput)

	currString := inputString;

	for i := 0; i < 40; i++ {
		currString = lookAndSayString(currString)
	}

	length := len(currString)

	return strconv.Itoa(length)
}

func part2(partInput string) string {
	inputString := aoc.ParseInput(partInput)

	currString := inputString;

	for i := 0; i < 50; i++ {
		currString = lookAndSayString(currString)
	}

	length := len(currString)

	return strconv.Itoa(length)
}

func lookAndSayString(look string) (say string) {
	var sayBuilder strings.Builder

	currCount := 0
	currRune := []rune(look)[0]

	for i, curr := range look {
		if curr == currRune {
			currCount++

			// If we're on the last character finish writing to the buffer
			if i == len(look) - 1 {
				sayBuilder.WriteString(strconv.Itoa(currCount))
				sayBuilder.WriteRune(currRune)
			}
		
			continue
		}

		sayBuilder.WriteString(strconv.Itoa(currCount))
		sayBuilder.WriteRune(currRune)

		currRune = curr
		currCount = 1

		// If we're on the last character finish writing to the buffer
		if i == len(look) - 1 {
			sayBuilder.WriteString(strconv.Itoa(currCount))
			sayBuilder.WriteRune(currRune)
		}
	}

	return sayBuilder.String()
}
