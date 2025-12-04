package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"andrewhaine/advent-of-go/util/gridutil"
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
	grid := gridutil.ParseGridFromInput(parsed, cellCallback)

	accessibleRollCount := 0

	for y, row := range grid {
		for x, cell := range row {
			if cell != "@" {
				continue
			}

			rollCount := 0
			for _, delta := range gridutil.AdjacentDeltas {
				checkCoord := gridutil.Coordinate{X: x + delta[0], Y: y + delta[1]}

				if !gridutil.WithinGrid(checkCoord, grid) {
					continue
				}

				if grid[checkCoord.Y][checkCoord.X] == "@" {
					rollCount++
				}
			}

			if rollCount < 4 {
				accessibleRollCount++
			}
		}
	}

	return strconv.Itoa(accessibleRollCount)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	grid := gridutil.ParseGridFromInput(parsed, cellCallback)

	removedCount := 0
	canRemove := true

	for canRemove {
		toRemove := []gridutil.Coordinate{}
		for y, row := range grid {
			for x, cell := range row {
				if cell != "@" {
					continue
				}

				rollCount := 0
				for _, delta := range gridutil.AdjacentDeltas {
					checkCoord := gridutil.Coordinate{X: x + delta[0], Y: y + delta[1]}

					if !gridutil.WithinGrid(checkCoord, grid) {
						continue
					}

					if grid[checkCoord.Y][checkCoord.X] == "@" {
						rollCount++
					}
				}

				if rollCount < 4 {
					toRemove = append(toRemove, gridutil.Coordinate{X: x, Y: y})
				}
			}
		}

		canRemove = len(toRemove) > 0

		if !canRemove {
			break
		}

		for _, coord := range toRemove {
			grid[coord.Y][coord.X] = "."
		}

		removedCount += len(toRemove)
	}

	return strconv.Itoa(removedCount)
}

func cellCallback(raw string) string {
	return raw
}
