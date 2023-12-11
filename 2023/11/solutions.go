package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"andrewhaine/advent-of-go/util/sliceutil"
	"math"
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

type Coordinate struct {
	x int
	y int
}

type Universe map[int][]string

func part1(partInput string) string {
	parsed := aoc.ParseInput(partInput)

	universe := generateUniverse(parsed)
	galaxies := findGalaxies(universe, 2)

	totalDistance := 0
	galaxiesMapped := map[Coordinate][]Coordinate{}

	for _, galaxy := range galaxies {
		distancesFromGalaxy := getDistancesTaxicab(galaxy, galaxies)

		for foundGalaxy, distance := range distancesFromGalaxy {
			if !slices.Contains(galaxiesMapped[galaxy], foundGalaxy) {
				totalDistance += distance
			}
			galaxiesMapped[foundGalaxy] = append(galaxiesMapped[foundGalaxy], galaxy)
		}
	}

	return strconv.Itoa(totalDistance)
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)

	universe := generateUniverse(parsed)
	galaxies := findGalaxies(universe, 1000000)

	totalDistance := 0
	galaxiesMapped := map[Coordinate][]Coordinate{}

	for _, galaxy := range galaxies {
		distancesFromGalaxy := getDistancesTaxicab(galaxy, galaxies)

		for foundGalaxy, distance := range distancesFromGalaxy {
			if !slices.Contains(galaxiesMapped[galaxy], foundGalaxy) {
				totalDistance += distance
			}
			galaxiesMapped[foundGalaxy] = append(galaxiesMapped[foundGalaxy], galaxy)
		}
	}

	return strconv.Itoa(totalDistance)
}

func generateUniverse(stringInput string) Universe {
	universe := Universe{}
	rows := strings.Split(stringInput, "\n")

	for i, row := range rows {
		universe[i] = strings.Split(row, "")
	}

	return universe
}

func findEmptyRowsAndColumns(universe Universe) (emptyRows []int, emptyColumns []int) {
	emptyRows = []int{}
	emptyColumns = []int{}

	for i, row := range universe {
		if !slices.Contains(row, "#") {
			emptyRows = append(emptyRows, i)
		}
	}

	for i := 0; i < len(universe[0]); i++ {
		column := []string{}
		for _, row := range universe {
			column = append(column, row[i])
		}
		empty := sliceutil.Every(column, func(index int, item string) bool {
			return item == "."
		})
		if empty {
			emptyColumns = append(emptyColumns, i)
		}
	}

	return
}

func findGalaxies(universe Universe, expansionFactor int) []Coordinate {
	galaxies := []Coordinate{}
	emptyRows, emptyColumns := findEmptyRowsAndColumns(universe)

	for i, row := range universe {
		previousEmptyRows := 0

		for _, emptyRow := range emptyRows {
			if i > emptyRow {
				previousEmptyRows += expansionFactor - 1
			}
		}

		for j, column := range row {
			previousEmptyColumns := 0
			for _, emptyColumn := range emptyColumns {
				if j > emptyColumn {
					previousEmptyColumns += expansionFactor - 1
				}
			}

			if column == "#" {
				galaxies = append(galaxies, Coordinate{j + previousEmptyColumns, i + previousEmptyRows})
			}
		}
	}

	return galaxies
}

func getDistancesTaxicab(galaxy Coordinate, galaxies []Coordinate) map[Coordinate]int {
	distances := map[Coordinate]int{}
	for _, toGalaxy := range galaxies {
		if galaxy == toGalaxy {
			continue
		}

		distanceX := math.Abs(float64(galaxy.x - toGalaxy.x))
		distanceY := math.Abs(float64(galaxy.y - toGalaxy.y))

		distances[toGalaxy] = int(distanceX + distanceY)
	}
	return distances
}

/* ============================================================================================================= */

// Soooo, this uses BFS which is going to take waaaay too long for our input

// type Node struct {
// 	coordinate Coordinate
// 	distance   int
// }

// func part1(partInput string) string {
// 	parsed := aoc.ParseInput(partInput)

// 	universe := generateUniverse(parsed)

// 	galaxies := findGalaxies(universe)

// 	totalDistance := 0
// 	galaxiesMapped := map[Coordinate][]Coordinate{}

// 	for _, galaxy := range galaxies {
// 		distancesFromGalaxy := getDistancesBFS(galaxy, galaxies, universe)

// 		for foundGalaxy, distance := range distancesFromGalaxy {
// 			if !slices.Contains(galaxiesMapped[galaxy], foundGalaxy) {
// 				totalDistance += distance
// 			}
// 			galaxiesMapped[foundGalaxy] = append(galaxiesMapped[foundGalaxy], galaxy)
// 		}
// 	}

// 	return strconv.Itoa(totalDistance)
// }

// func getDistancesBFS(galaxy Coordinate, galaxies []Coordinate, universe Universe) map[Coordinate]int {
// 	distances := map[Coordinate]int{}

// 	startNode := Node{galaxy, 0}
// 	visited := map[Coordinate]bool{}
// 	queue := []Node{startNode}

// 	for len(queue) > 0 && len(distances) < len(galaxies) {
// 		current := queue[0]
// 		queue = queue[1:]

// 		if visited[current.coordinate] {
// 			continue
// 		}

// 		visited[current.coordinate] = true
// 		locationType := universe[current.coordinate.y][current.coordinate.x]

// 		if locationType == "#" {
// 			distances[current.coordinate] = current.distance
// 		}

// 		queue = append(queue, getNeighbours(current, universe)...)
// 	}

// 	return distances
// }

// func getNeighbours(node Node, universe Universe) []Node {
// 	emptyRows, emptyColumns := findEmptyRowsAndColumns(universe)

// 	directions := [][2]int{
// 		{0, 1},
// 		{1, 0},
// 		{0, -1},
// 		{-1, 0},
// 	}

// 	nodeDistance := node.distance
// 	neighbours := []Node{}

// 	for _, delta := range directions {
// 		newX := node.coordinate.x + delta[0]
// 		newY := node.coordinate.y + delta[1]

// 		if newY == -1 || newX == -1 || newY == len(universe) || newX == len(universe[node.coordinate.y]) {
// 			continue
// 		}

// 		neighbourCoord := Coordinate{newX, newY}
// 		neighbourDist := nodeDistance + 1

// 		if newX != node.coordinate.x && slices.Contains(emptyColumns, newX) {
// 			neighbourDist += 1
// 		}

// 		if newY != node.coordinate.y && slices.Contains(emptyRows, newY) {
// 			neighbourDist += 1
// 		}

// 		neighbours = append(neighbours, Node{neighbourCoord, neighbourDist})
// 	}

// 	return neighbours
// }
