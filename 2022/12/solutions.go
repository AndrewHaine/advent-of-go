package main

import (
	"andrewhaine/advent-of-go/util/aoc"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

var alphabet = "abcdefghijklmnopqrstuvwxyz"

type Coordinate struct {
	x int
	y int
}

type Node struct {
	coordinate Coordinate
	distance   int
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
	terrain := map[int]string{}

	for i, row := range strings.Split(parsed, "\n") {
		terrain[i] = row
	}

	startCoords := getStartCoords(terrain)
	startNode := Node{startCoords, 0}
	visited := map[Coordinate]bool{}
	queue := []Node{startNode}

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if visited[c.coordinate] {
			continue
		}

		visited[c.coordinate] = true
		level := string(terrain[c.coordinate.y][c.coordinate.x])

		if level == "S" {
			return strconv.Itoa(c.distance)
		}

		queue = append(queue, getNeighbours(c, terrain)...)
	}

	return "-1"
}

func part2(partInput string) string {
	parsed := aoc.ParseInput(partInput)
	terrain := map[int]string{}

	for i, row := range strings.Split(parsed, "\n") {
		terrain[i] = row
	}

	startCoords := getStartCoords(terrain)
	startNode := Node{startCoords, 0}
	visited := map[Coordinate]bool{}
	queue := []Node{startNode}

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if visited[c.coordinate] {
			continue
		}

		visited[c.coordinate] = true
		level := string(terrain[c.coordinate.y][c.coordinate.x])

		if level == "S" || level == "a" {
			return strconv.Itoa(c.distance)
		}

		queue = append(queue, getNeighbours(c, terrain)...)
	}

	return "-1"
}

func getStartCoords(terrain map[int]string) Coordinate {
	startCoords := Coordinate{0, 0}

	for i, row := range terrain {
		startIndex := strings.Index(row, "E")

		if startIndex > -1 {
			startCoords.x = startIndex
			startCoords.y = i
		}
	}

	return startCoords
}

func getNeighbours(node Node, terrain map[int]string) []Node {
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	nodeLevel := string(terrain[node.coordinate.y][node.coordinate.x])
	nodeDistance := node.distance
	neighbours := []Node{}

	for _, delta := range directions {
		newY := node.coordinate.y + delta[1]
		newX := node.coordinate.x + delta[0]

		if newY == -1 || newX == -1 || newY == len(terrain) || newX == len(terrain[node.coordinate.y]) {
			continue
		}

		neighbourLevel := string(terrain[newY][newX])
		heightDiff := getHeightDifference(nodeLevel, neighbourLevel)

		if heightDiff > 1 {
			continue
		}

		neighbourCoord := Coordinate{newX, newY}
		neighbourDist := nodeDistance + 1

		neighbours = append(neighbours, Node{neighbourCoord, neighbourDist})
	}

	return neighbours
}

func getHeightDifference(a string, b string) int {
	if a == "S" {
		a = "a"
	}

	if a == "E" {
		a = "z"
	}

	if b == "S" {
		b = "a"
	}

	if b == "E" {
		b = "z"
	}

	return strings.Index(alphabet, a) - strings.Index(alphabet, b)
}
