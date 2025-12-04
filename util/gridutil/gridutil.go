package gridutil

import "strings"

type Coordinate struct {
	X int
	Y int
}

type Grid[C comparable] [][]C

func ParseGridFromInput[C comparable](input string, cellCb func(cell string) C) Grid[C] {
	grid := Grid[C]{}
	for _, rowRaw := range strings.Split(input, "\n") {
		row := []C{}
		for _, colRaw := range strings.Split(rowRaw, "") {
			row = append(row, cellCb(colRaw))
		}
		grid = append(grid, row)
	}
	return grid
}

func WithinGrid[C comparable](coord Coordinate, grid Grid[C]) bool {
	return (coord.Y > -1 &&
		coord.X > -1 &&
		coord.Y < len(grid) &&
		coord.X < len(grid[coord.Y]))
}

var AdjacentDeltas = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{-1, -1},
	{-1, 1},
	{1, 1},
	{1, -1},
}
