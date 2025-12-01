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

// Return the number of times the dial ends on zero after moving.
func part1(partInput string) string {
	instructions := strings.Split(aoc.ParseInput(partInput), "\n")
	instReg := regexp.MustCompile(`(L|R)(\d+)`)

	zeroCount := 0
	val := 50

	for _, inst := range instructions {
		matches := instReg.FindStringSubmatch(inst)
		dir := matches[1]
		length, _ := strconv.Atoi(matches[2])

		if dir == "L" {
			val -= length
		} else {
			val += length
		}

		if isZero(val) {
			zeroCount++
		}
	}

	return strconv.Itoa(zeroCount)
}

// Return the number of times the dial ends on zero during or after moving the dial.
func part2(partInput string) string {
	instructions := strings.Split(aoc.ParseInput(partInput), "\n")
	instReg := regexp.MustCompile(`(L|R)(\d+)`)

	zeroCount := 0
	val := 50

	for _, inst := range instructions {
		matches := instReg.FindStringSubmatch(inst)
		dir := matches[1]
		length, _ := strconv.Atoi(matches[2])

		for length > 0 {
			if dir == "L" {
				val -= 1
			} else {
				val += 1
			}

			length -= 1

			if isZero(val) {
				zeroCount++
			}
		}
	}

	return strconv.Itoa(zeroCount)
}

/**
 * Pass in the current value (not wrapped - i.e. this can be greater than 99 and less than zero)
 * we know that if it is divisible by 50 and the factor is odd we have hit zero
 * i.e. 150 / 50 = 3 (isZero) 200 / 50 = 4 (isNotZero)
 */
func isZero(val int) bool {
	return val%100 == 0
}
