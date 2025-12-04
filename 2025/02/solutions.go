package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"regexp"
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
	ranges := strings.Split(parsed, ",")
	tot := calculateTotal(ranges, isInvalidPart1)

	return strconv.Itoa(tot)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	ranges := strings.Split(parsed, ",")
	tot := calculateTotal(ranges, isInvalidPart2)

	return strconv.Itoa(tot)
}

func isInvalidPart1(id int) bool {
	idStr := strconv.Itoa(id)
	idStrLen := len(idStr)

	// Odd-length IDs cannot be invalid
	if idStrLen%2 != 0 {
		return false
	}

	return idStr[:idStrLen/2] == idStr[idStrLen/2:]
}

func isInvalidPart2(id int) bool {
	idStr := strconv.Itoa(id)

	possibleLens := []int{}

	for i := range len(idStr) {
		if i == 0 {
			continue
		}

		if len(idStr)%(i) == 0 {
			possibleLens = append(possibleLens, i)
		}
	}

	isInvalid := false

	// Go from large -> small chunks to save time
	slices.Reverse(possibleLens)

	for _, length := range possibleLens {
		start := 0
		end := length
		chunks := []string{}

		for end <= len(idStr) {
			chunks = append(chunks, idStr[start:end])
			end += length
			start += length
		}

		chunksAreEqual := true
		for _, chunk := range chunks[1:] {
			if chunk != chunks[0] {
				chunksAreEqual = false
				break
			}
		}

		if chunksAreEqual {
			isInvalid = true
			break
		}
	}

	return isInvalid
}

func calculateTotal(ranges []string, isInvalidCb func(id int) bool) int {
	rangeReg := regexp.MustCompile(`(\d+)-(\d+)`)

	invalidIds := []int{}

	for _, rang := range ranges {
		matches := rangeReg.FindStringSubmatch(rang)
		start, _ := strconv.Atoi(matches[1])
		end, _ := strconv.Atoi(matches[2])

		curr := start

		for curr <= end {
			if isInvalidCb(curr) {
				invalidIds = append(invalidIds, curr)
			}
			curr++
		}
	}

	tot := 0

	for _, invalidId := range invalidIds {
		tot += invalidId
	}

	return tot
}
