package aoc

import (
	"fmt"
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
