package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"andrewhaine/advent-of-go/util/maputil"
	"bytes"
	"strconv"
	"strings"
	"unicode"

	_ "embed"
)

//go:embed input.txt
var input string

var numberNames = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

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
	lines := strings.Split(parsed, "\n")
	total := 0

	for _, line := range lines {
		digits := []rune{}
		for _, char := range line {
			if unicode.IsNumber(char) {
				digits = append(digits, char)
			}
		}

		total += concatRunesToInt(digits)
	}

	return strconv.Itoa(total)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	lines := strings.Split(parsed, "\n")
	total := 0

	for _, line := range lines {
		digitMap := map[int]rune{}

		for index, char := range line {
			if unicode.IsNumber(char) {
				digitMap[index] = char
			}
		}

		for numberName, value := range numberNames {
			firstIndex := strings.Index(line, numberName)
			lastIndex := strings.LastIndex(line, numberName)
			if firstIndex > -1 {
				digitMap[firstIndex] = value
			}

			if lastIndex > -1 && lastIndex != firstIndex {
				digitMap[lastIndex] = value
			}
		}

		digits := maputil.SortedValues(digitMap)

		total += concatRunesToInt(digits)
	}

	return strconv.Itoa(total)
}

func concatRunesToInt(runes []rune) int {
	var stringVal bytes.Buffer
	stringVal.WriteRune(runes[0])
	stringVal.WriteRune(runes[len(runes)-1])
	intVal, _ := strconv.Atoi(stringVal.String())
	return intVal
}
