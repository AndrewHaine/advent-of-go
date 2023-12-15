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

type Lens struct {
	label       string
	focalLength int
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
	sequence := strings.Split(parsed, ",")

	total := 0

	for _, step := range sequence {
		total += runHash(step)
	}

	return strconv.Itoa(total)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	sequence := strings.Split(parsed, ",")

	boxes := map[int][]Lens{}

	for _, step := range sequence {
		if strings.Contains(step, "=") {
			split := strings.Split(step, "=")
			label, focalLengthString := split[0], split[1]
			focalLength, _ := strconv.Atoi(focalLengthString)

			box := runHash(label)
			lens := Lens{label, focalLength}

			if boxContents, ok := boxes[box]; ok {
				foundIndex := slices.IndexFunc(boxContents, func(boxLens Lens) bool {
					return boxLens.label == label
				})

				if foundIndex > -1 {
					boxes[box] = slices.Replace(boxContents, foundIndex, foundIndex+1, lens)
				} else {
					boxes[box] = append(boxContents, lens)
				}

			} else {
				boxes[box] = []Lens{lens}
			}
		}

		if strings.Contains(step, "-") {
			split := strings.Split(step, "-")
			label := split[0]

			box := runHash(label)

			if boxContents, ok := boxes[box]; ok {
				boxes[box] = slices.DeleteFunc(boxContents, func(boxLens Lens) bool {
					return boxLens.label == label
				})
			}
		}
	}

	power := 0

	for box, boxLenses := range boxes {
		for i, lens := range boxLenses {
			power += (1 + box) * (1 + i) * lens.focalLength
		}
	}

	return strconv.Itoa(power)
}

func runHash(step string) int {
	current := 0

	for _, char := range step {
		current += int(char)
		current *= 17
		current = current % 256
	}

	return current
}
