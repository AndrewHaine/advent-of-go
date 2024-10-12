# 🎄 Advent of Go 🎄

Advent of Code solutions written in Go.

## Usage

### Creating a new day

This package comes with template files for each day, simply run `make day` and follow the instructions to generate the relevant files.

### Testing solutions

Each AoC puzzle has an example input along with the expected outcome, in order to test your solution paste the provided input into `input_test.txt` for the day, cd into the day directory and run `go test -v`.

_Note: Some puzzles provide multiple test inputs, Add multiple tests/input files as required._

### Running solutions

Getting the answer for a solution is as simple as changing into the day directory and running:

```bash
go run . -part=1
```

## A note on puzzle inputs

Your puzzle input should be pasted into the `input.txt` file for each day.

Puzzle inputs are not included in this repository, this is a request from the AoC team.
