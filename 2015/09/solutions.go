package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"math"
	"regexp"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type LocationDistanceMap map[string]int

type LocationMatrix map[string]LocationDistanceMap

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
	distances := strings.Split(parsed, "\n")

	matrix := generateLocationMatrixFromInput(distances)
	locations := []string{}
	shortestDistance := math.MaxInt

	for location := range matrix {
		locations = append(locations, location)
	}

	allPermutations := [][]string{}

	getHeapPermutations(locations, len(locations), &allPermutations)

	for _, permutation := range allPermutations {
		distanceForPermutation := 0
		for i, location := range permutation {
			if i == len(permutation)-1 {
				break
			}

			distanceForPermutation += matrix[location][permutation[i+1]]
		}

		if distanceForPermutation < shortestDistance {
			shortestDistance = distanceForPermutation
		}
	}

	return strconv.Itoa(shortestDistance)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	distances := strings.Split(parsed, "\n")

	matrix := generateLocationMatrixFromInput(distances)
	locations := []string{}
	longestDistance := 0

	for location := range matrix {
		locations = append(locations, location)
	}

	allPermutations := [][]string{}

	getHeapPermutations(locations, len(locations), &allPermutations)

	for _, permutation := range allPermutations {
		distanceForPermutation := 0
		for i, location := range permutation {
			if i == len(permutation)-1 {
				break
			}

			distanceForPermutation += matrix[location][permutation[i+1]]
		}

		if distanceForPermutation > longestDistance {
			longestDistance = distanceForPermutation
		}
	}

	return strconv.Itoa(longestDistance)
}

func generateLocationMatrixFromInput(inputDistances []string) LocationMatrix {
	matrix := LocationMatrix{}

	inputRegex := regexp.MustCompile(`([a-zA-Z]+) to ([a-zA-Z]+) = (\d+)`)

	for _, inputString := range inputDistances {
		parts := inputRegex.FindStringSubmatch(inputString)

		from, to, distanceString := parts[1], parts[2], parts[3]
		distance, _ := strconv.Atoi(distanceString)

		if _, ok := matrix[from]; !ok {
			locationDistance := LocationDistanceMap{to: distance}
			matrix[from] = locationDistance
		} else {
			matrix[from][to] = distance
		}

		if _, ok := matrix[to]; !ok {
			locationDistance := LocationDistanceMap{from: distance}
			matrix[to] = locationDistance
		} else {
			matrix[to][from] = distance
		}
	}

	return matrix
}

func getHeapPermutations(values []string, size int, allPermutations *[][]string) {
	if size == 1 {
		perm := make([]string, len(values))
		copy(perm, values)
		*allPermutations = append(*allPermutations, perm)
	}

	for i := 0; i < size; i++ {
		getHeapPermutations(values, size-1, allPermutations)

		if size%2 == 1 {
			values[0], values[size-1] = values[size-1], values[0]
		} else {
			values[i], values[size-1] = values[size-1], values[i]
		}
	}
}
