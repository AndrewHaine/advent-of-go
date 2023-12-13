package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"reflect"
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

	blocks := strings.Split(parsed, "\n\n")
	total := 0

	for _, block := range blocks {
		reflectedRows, reflectedColumns := getReflectedRowOrColumnForBlock(block, false)

		if len(reflectedRows) > 0 {
			total += reflectedRows[0] * 100
			continue
		}

		total += reflectedColumns[0]
	}

	return strconv.Itoa(total)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)

	blocks := strings.Split(parsed, "\n\n")
	total := 0

	for _, block := range blocks {
		reflectedRows, reflectedColumns := getReflectedRowOrColumnForBlock(block, false)
		reflectedRowsFixed, reflectedColumnsFixed := getReflectedRowOrColumnForBlock(block, true)

		// We now have two reflected rows - pick the one that doesn't exist in the original data
		if len(reflectedRowsFixed) > 1 {
			reflectedRow := 0
			for _, row := range reflectedRowsFixed {
				if !slices.Contains(reflectedRows, row) {
					reflectedRow = row
				}
			}
			total += reflectedRow * 100
			continue
		}

		// We now have two reflected columns - pick the one that doesn't exist
		if len(reflectedColumnsFixed) > 1 {
			reflectedColumn := 0
			for _, column := range reflectedColumnsFixed {
				if !slices.Contains(reflectedColumns, column) {
					reflectedColumn = column
				}
			}
			total += reflectedColumn
			continue
		}

		// In this case we must have one of each? So pick the opposite
		if len(reflectedRows) > 0 {
			total += reflectedColumnsFixed[0]
			continue
		}

		total += 100 * reflectedRowsFixed[0]

	}

	return strconv.Itoa(total)
}

func getReflectedRowOrColumnForBlock(block string, fixSmudge bool) (reflectedRows []int, reflectedColumns []int) {
	blockGrid := strings.Split(block, "\n")

	for i := 0; i < len(blockGrid)-1; i++ {
		smudgeFixed := false

		// If the diff is equal to 1 fix the smudge
		topDiffCount := 0

		for j, pattern := range blockGrid[i] {
			if blockGrid[i+1][j] != byte(pattern) {
				topDiffCount += 1
			}
		}

		if topDiffCount == 1 && fixSmudge {
			smudgeFixed = true
		}

		// Each char in row matches the char in the next row
		if reflect.DeepEqual(blockGrid[i], blockGrid[i+1]) || (fixSmudge && smudgeFixed) {
			reflectionSize := 1
			reflecting := true
			perfectReflection := false

			// Expand the search up/down to see the size of the reflection
			for reflecting {
				top := i - reflectionSize
				bottom := i + 1 + reflectionSize

				if bottom > len(blockGrid)-1 || top < 0 {
					reflecting = false
					perfectReflection = true
					break
				}

				diffCount := 0

				for j, pattern := range blockGrid[top] {
					if blockGrid[bottom][j] != byte(pattern) {
						diffCount += 1
					}
				}

				if reflect.DeepEqual(blockGrid[top], blockGrid[bottom]) || (fixSmudge && diffCount == 1 && !smudgeFixed) {
					reflectionSize += 1
					if diffCount == 1 {
						smudgeFixed = true
					}
				} else {
					reflecting = false
					break
				}

			}

			if perfectReflection {
				reflectedRows = append(reflectedRows, i+1)
			}
		}
	}

	for i := 0; i < len(blockGrid[0])-1; i++ {
		col1 := []string{}
		col2 := []string{}

		for _, row := range blockGrid {
			col1 = append(col1, string(row[i]))
			col2 = append(col2, string(row[i+1]))
		}

		smudgeFixed := false

		// If the diff is equal to 1 fix the smudge
		topDiffCount := 0

		for j, pattern := range col1 {
			if col2[j] != pattern {
				topDiffCount += 1
			}
		}

		if topDiffCount == 1 && fixSmudge {
			smudgeFixed = true
		}

		if reflect.DeepEqual(col1, col2) || (fixSmudge && smudgeFixed) {
			reflectionSize := 1
			reflecting := true
			perfectReflection := false

			// Expand the search left/right to see the size of the reflection
			for reflecting {
				left := i - reflectionSize
				right := i + 1 + reflectionSize

				if right > len(blockGrid[0])-1 || left < 0 {
					reflecting = false
					perfectReflection = true
					break
				}

				col1 := []string{}
				col2 := []string{}

				for _, row := range blockGrid {
					col1 = append(col1, string(row[left]))
					col2 = append(col2, string(row[right]))
				}

				diffCount := 0

				for j, pattern := range col1 {
					if col2[j] != pattern {
						diffCount += 1
					}
				}

				if reflect.DeepEqual(col1, col2) || (fixSmudge && diffCount == 1 && !smudgeFixed) {
					reflectionSize += 1
					if diffCount == 1 {
						smudgeFixed = true
					}
				} else {
					reflecting = false
					break
				}
			}

			if perfectReflection {
				reflectedColumns = append(reflectedColumns, i+1)
			}
		}
	}

	return
}
