package aoc

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

func PrintSolution(output string) {
	ansStyle := color.New(color.FgHiYellow, color.Bold).SprintFunc()
	fmt.Printf("The answer is: %s\n", ansStyle(output))
}

func ParseInput(input string) string {
	trimmed := strings.Trim(input, "\n")

	if len(trimmed) == 0 {
		panic("Empty input")
	}

	return trimmed
}

func PartFlag() int {
	var part int
	flag.IntVar(&part, "part", 1, "Part 1 or 2")
	flag.Parse()

	if part != 1 && part != 2 {
		log.Fatalf("Part %d is out of range", part)
	}

	return part
}
